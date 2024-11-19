// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

// Package main contains mqtt-adapter main function to start the mqtt-adapter service.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	chclient "github.com/absmach/callhome/pkg/client"
	"github.com/absmach/magistrala"
	mglog "github.com/absmach/magistrala/logger"
	"github.com/absmach/magistrala/mqtt"
	"github.com/absmach/magistrala/mqtt/events"
	mqtttracing "github.com/absmach/magistrala/mqtt/tracing"
	"github.com/absmach/magistrala/pkg/errors"
	"github.com/absmach/magistrala/pkg/grpcclient"
	jaegerclient "github.com/absmach/magistrala/pkg/jaeger"
	"github.com/absmach/magistrala/pkg/messaging/brokers"
	brokerstracing "github.com/absmach/magistrala/pkg/messaging/brokers/tracing"
	"github.com/absmach/magistrala/pkg/messaging/handler"
	mqttpub "github.com/absmach/magistrala/pkg/messaging/mqtt"
	"github.com/absmach/magistrala/pkg/server"
	"github.com/absmach/magistrala/pkg/uuid"
	mgate "github.com/absmach/mgate"
	mgatemqtt "github.com/absmach/mgate/pkg/mqtt"
	"github.com/absmach/mgate/pkg/mqtt/websocket"
	"github.com/absmach/mgate/pkg/session"
	"github.com/caarlos0/env/v11"
	"github.com/cenkalti/backoff/v4"
	"golang.org/x/sync/errgroup"
)

const (
	svcName         = "mqtt"
	envPrefixThings = "MG_THINGS_AUTH_GRPC_"
	wsPathPrefix    = "/mqtt"
)

type config struct {
	LogLevel              string        `env:"MG_MQTT_ADAPTER_LOG_LEVEL"                    envDefault:"info"`
	MQTTPort              string        `env:"MG_MQTT_ADAPTER_MQTT_PORT"                    envDefault:"1883"`
	MQTTTargetHost        string        `env:"MG_MQTT_ADAPTER_MQTT_TARGET_HOST"             envDefault:"localhost"`
	MQTTTargetPort        string        `env:"MG_MQTT_ADAPTER_MQTT_TARGET_PORT"             envDefault:"1883"`
	MQTTForwarderTimeout  time.Duration `env:"MG_MQTT_ADAPTER_FORWARDER_TIMEOUT"            envDefault:"30s"`
	MQTTTargetHealthCheck string        `env:"MG_MQTT_ADAPTER_MQTT_TARGET_HEALTH_CHECK"     envDefault:""`
	MQTTQoS               uint8         `env:"MG_MQTT_ADAPTER_MQTT_QOS"                     envDefault:"1"`
	HTTPPort              string        `env:"MG_MQTT_ADAPTER_WS_PORT"                      envDefault:"8080"`
	HTTPTargetHost        string        `env:"MG_MQTT_ADAPTER_WS_TARGET_HOST"               envDefault:"localhost"`
	HTTPTargetPort        string        `env:"MG_MQTT_ADAPTER_WS_TARGET_PORT"               envDefault:"8080"`
	HTTPTargetPath        string        `env:"MG_MQTT_ADAPTER_WS_TARGET_PATH"               envDefault:"/mqtt"`
	Instance              string        `env:"MG_MQTT_ADAPTER_INSTANCE"                     envDefault:""`
	JaegerURL             url.URL       `env:"MG_JAEGER_URL"                                envDefault:"http://localhost:4318/v1/traces"`
	BrokerURL             string        `env:"MG_MESSAGE_BROKER_URL"                        envDefault:"nats://localhost:4222"`
	SendTelemetry         bool          `env:"MG_SEND_TELEMETRY"                            envDefault:"true"`
	InstanceID            string        `env:"MG_MQTT_ADAPTER_INSTANCE_ID"                  envDefault:""`
	ESURL                 string        `env:"MG_ES_URL"                                    envDefault:"nats://localhost:4222"`
	TraceRatio            float64       `env:"MG_JAEGER_TRACE_RATIO"                        envDefault:"1.0"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to load %s configuration : %s", svcName, err)
	}

	logger, err := mglog.New(os.Stdout, cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to init logger: %s", err.Error())
	}

	var exitCode int
	defer mglog.ExitWithError(&exitCode)

	if cfg.InstanceID == "" {
		if cfg.InstanceID, err = uuid.New().ID(); err != nil {
			logger.Error(fmt.Sprintf("failed to generate instanceID: %s", err))
			exitCode = 1
			return
		}
	}

	if cfg.MQTTTargetHealthCheck != "" {
		notify := func(e error, next time.Duration) {
			logger.Info(fmt.Sprintf("Broker not ready: %s, next try in %s", e.Error(), next))
		}

		err := backoff.RetryNotify(healthcheck(cfg), backoff.NewExponentialBackOff(), notify)
		if err != nil {
			logger.Error(fmt.Sprintf("MQTT healthcheck limit exceeded, exiting. %s ", err))
			exitCode = 1
			return
		}
	}

	serverConfig := server.Config{
		Host: cfg.HTTPTargetHost,
		Port: cfg.HTTPTargetPort,
	}

	tp, err := jaegerclient.NewProvider(ctx, svcName, cfg.JaegerURL, cfg.InstanceID, cfg.TraceRatio)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to init Jaeger: %s", err))
		exitCode = 1
		return
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			logger.Error(fmt.Sprintf("Error shutting down tracer provider: %v", err))
		}
	}()
	tracer := tp.Tracer(svcName)

	bsub, err := brokers.NewPubSub(ctx, cfg.BrokerURL, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to message broker: %s", err))
		exitCode = 1
		return
	}
	defer bsub.Close()
	bsub = brokerstracing.NewPubSub(serverConfig, tracer, bsub)

	mpub, err := mqttpub.NewPublisher(fmt.Sprintf("mqtt://%s:%s", cfg.MQTTTargetHost, cfg.MQTTTargetPort), cfg.MQTTQoS, cfg.MQTTForwarderTimeout)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create MQTT publisher: %s", err))
		exitCode = 1
		return
	}
	defer mpub.Close()

	fwd := mqtt.NewForwarder(brokers.SubjectAllChannels, logger)
	fwd = mqtttracing.New(serverConfig, tracer, fwd, brokers.SubjectAllChannels)
	if err := fwd.Forward(ctx, svcName, bsub, mpub); err != nil {
		logger.Error(fmt.Sprintf("failed to forward message broker messages: %s", err))
		exitCode = 1
		return
	}

	np, err := brokers.NewPublisher(ctx, cfg.BrokerURL)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to message broker: %s", err))
		exitCode = 1
		return
	}
	defer np.Close()
	np = brokerstracing.NewPublisher(serverConfig, tracer, np)

	es, err := events.NewEventStore(ctx, cfg.ESURL, cfg.Instance)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create %s event store : %s", svcName, err))
		exitCode = 1
		return
	}

	thingsClientCfg := grpcclient.Config{}
	if err := env.ParseWithOptions(&thingsClientCfg, env.Options{Prefix: envPrefixThings}); err != nil {
		logger.Error(fmt.Sprintf("failed to load %s auth configuration : %s", svcName, err))
		exitCode = 1
		return
	}

	thingsClient, thingsHandler, err := grpcclient.SetupThingsClient(ctx, thingsClientCfg)
	if err != nil {
		logger.Error(err.Error())
		exitCode = 1
		return
	}
	defer thingsHandler.Close()

	logger.Info("Things service gRPC client successfully connected to things gRPC server " + thingsHandler.Secure())

	h := mqtt.NewHandler(np, es, logger, thingsClient)
	h = handler.NewTracing(tracer, h)

	if cfg.SendTelemetry {
		chc := chclient.New(svcName, magistrala.Version, logger, cancel)
		go chc.CallHome(ctx)
	}

	var interceptor session.Interceptor
	logger.Info(fmt.Sprintf("Starting MQTT proxy on port %s", cfg.MQTTPort))
	g.Go(func() error {
		return proxyMQTT(ctx, cfg, logger, h, interceptor)
	})

	logger.Info(fmt.Sprintf("Starting MQTT over WS  proxy on port %s", cfg.HTTPPort))
	g.Go(func() error {
		return proxyWS(ctx, cfg, logger, h, interceptor)
	})

	g.Go(func() error {
		return stopSignalHandler(ctx, cancel, logger)
	})

	if err := g.Wait(); err != nil {
		logger.Error(fmt.Sprintf("mProxy terminated: %s", err))
	}
}

func proxyMQTT(ctx context.Context, cfg config, logger *slog.Logger, sessionHandler session.Handler, interceptor session.Interceptor) error {
	config := mgate.Config{
		Address: fmt.Sprintf(":%s", cfg.MQTTPort),
		Target:  fmt.Sprintf("%s:%s", cfg.MQTTTargetHost, cfg.MQTTTargetPort),
	}
	mproxy := mgatemqtt.New(config, sessionHandler, interceptor, logger)

	errCh := make(chan error)
	go func() {
		errCh <- mproxy.Listen(ctx)
	}()

	select {
	case <-ctx.Done():
		logger.Info(fmt.Sprintf("proxy MQTT shutdown at %s", config.Target))
		return nil
	case err := <-errCh:
		return err
	}
}

func proxyWS(ctx context.Context, cfg config, logger *slog.Logger, sessionHandler session.Handler, interceptor session.Interceptor) error {
	config := mgate.Config{
		Address:    fmt.Sprintf("%s:%s", "", cfg.HTTPPort),
		Target:     fmt.Sprintf("ws://%s:%s%s", cfg.HTTPTargetHost, cfg.HTTPTargetPort, wsPathPrefix),
		PathPrefix: wsPathPrefix,
	}

	wp := websocket.New(config, sessionHandler, interceptor, logger)
	http.HandleFunc(wsPathPrefix, wp.ServeHTTP)

	errCh := make(chan error)

	go func() {
		errCh <- wp.Listen(ctx)
	}()

	select {
	case <-ctx.Done():
		logger.Info(fmt.Sprintf("proxy MQTT WS shutdown at %s", config.Target))
		return nil
	case err := <-errCh:
		return err
	}
}

func healthcheck(cfg config) func() error {
	return func() error {
		res, err := http.Get(cfg.MQTTTargetHealthCheck)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return errors.New(string(body))
		}
		return nil
	}
}

func stopSignalHandler(ctx context.Context, cancel context.CancelFunc, logger *slog.Logger) error {
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGABRT)
	select {
	case sig := <-c:
		defer cancel()
		logger.Info(fmt.Sprintf("%s service shutdown by signal: %s", svcName, sig))
		return nil
	case <-ctx.Done():
		return nil
	}
}
