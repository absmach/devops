# Mainflux Helm Chart

Helm Chart for the Mainflux IoT Platform

## Prerequisites

- Helm v3
- Stable Helm repo installed (`helm repo add stable https://kubernetes-charts.storage.googleapis.com/`)
- Nginx Ingress Controller
- If using the mTLS setup:
  - Certificate/Key installed as **TLS secret**
  - Intermediate certificate installed as **Generic secret**

## Configuration

The following table lists the configurable parameters and their default values.

| Parameter                            | Description                                                                | Default      |
| ------------------------------------ | -------------------------------------------------------------------------- | ------------ |
| defaults.logLevel                    | Log level                                                                  | debug        |
| defaults.image.pullPolicy            | Docker Image Pull Policy                                                   | IfNotPresent |
| defaults.image.repository            | Docker Image Repository                                                    | mainflux     |
| defaults.image.tag                   | Docker Image Tag                                                           | 0.11.0       |
| defaults.replicaCount                | Replicas of MQTT adapter, Things, Envoy and Authn                          | 3            |
| defaults.natsPort                    | NATS port                                                                  | 4222         |
| defaults.redisPort                   | Redis port                                                                 | 6379         |
| defaults.jaegerPort                  | Jaeger port                                                                | 6831         |
| nginx_internal.mtls.tls              | TLS secret which contains the server cert/key                              |              |
| nginx_internal.mtls.intermediate_crt | Generic secret which contains the intermediate cert used to verify clients |              |
| ingress.enabled                      | Should the Nginx Ingress be created                                        | true         |
| ingress.hostname                     | Hostname for the Nginx Ingress                                             |              |
| ingress.tls.hostname                 | Hostname of the Nginx Ingress certificate                                  |              |
| ingress.tls.secret                   | TLS secret for the Nginx Ingress                                           |              |
| nats.maxPayload                      | Maximum payload size in bytes that the NATS server will accept             | 268435456    |
| nats.replicaCount                    | NATS replicas                                                              | 3            |
| authn.dbPort                         | AuthN service DB port                                                      | 5432         |
| authn.grpcPort                       | AuthN service gRPC port                                                    | 8181         |
| authn.httpPort                       | AuthN service HTTP port                                                    | 8189         |
| users.dbPort                         | Users service DB port                                                      | 5432         |
| users.httpPort                       | Users service HTTP port                                                    | 8180         |
| things.dbPort                        | Things service DB port                                                     | 5432         |
| things.httpPort                      | Things service HTTP port                                                   | 8182         |
| things.authGrpcPort                  | Things service Auth gRPC port                                              | 8183         |
| things.authHttpPort                  | Things service Auth HTTP port                                              | 8989         |
| adapter_http.port                    | HTTP Adapter port                                                          | 8185         |
| mqtt.proxy.mqttPort                  | MQTT proxy port                                                            | 1884         |
| mqtt.proxy.wsPort                    | MQTT proxy WS port                                                         | 8081         |
| mqtt.broker.mqttPort                 | MQTT broker port                                                           | 1883         |
| mqtt.broker.wsPort                   | MQTT broker WS port                                                        | 8080         |
| mqtt.broker.persistentVolume.size    | MQTT broker data Persistent Volume size                                    | 5Gi          |
| adapter_coap.port                    | CoAP Adapter port                                                          | 5683         |
| ui.port                              | UI port                                                                    | 3000         |
| bootstrap.enabled                    | Enable bootstrap service                                                   | false        |
| bootstrap.dbPort                     | Bootstrap service DB port                                                  | 5432         |
| bootstrap.httpPort                   | Bootstrap service HTTP port                                                | 8182         |
| influxdb.enabled                     | Enable InfluxDB reader & writer                                            | false        |
| influxdb.dbPort                      | InfluxDB port                                                              | 8086         |
| influxdb.writer.httpPort             | InfluxDB writer HTTP port                                                  | 8900         |
| influxdb.reader.httpPort             | InfluxDB reader HTTP port                                                  | 8905         |
| adapter_opcua.enabled                | Enable OPC-UA Adapter                                                      | false        |
| adapter_opcua.httpPort               | OPC-UA Adapter HTTP port                                                   | 8188         |
| adapter_lora.enabled                 | Enable LoRa Adapter                                                        | false        |
| adapter_lora.httpPort                | LoRa Adapter HTTP port                                                     | 8187         |
| twins.enabled                        | Enable twins service                                                       | false        |
| twins.dbPort                         | Twins service DB port                                                      | 27017        |
| twins.httpPort                       | Twins service HTTP port                                                    | 9021         |

All Mainflux services (both core and add-ons) can have their `logLevel`, `image.pullPolicy`, `image.repository` and `image.tag` overridden.

Mainflux Core is a minimalistic set of required Mainflux services. They are all installed by default:

- authn
- users
- things
- adapter_http
- adapter_mqtt
- adapter_coap
- ui

Mainflux Add-ons are optional services that are disabled by default. Find in Configuration table paramaters for enabling them, i.e. to enable influxdb reader & writer you shoud run `helm install` with `--set influxdb=true`.
List of add-ons services in charts:

- bootstrap
- influxdb.writer
- influxdb.reader
- adapter_opcua
- adapter_lora
- twins

By default scale of MQTT adapter, Things, Envoy, Authn and NATS will be set to 3. It's recommended that you set this values to number of your nodes in Kubernetes cluster, i.e. `--set defaults.replicaCount=3 --set nats.replicaCount=3`

**Note:** make sure you run `helm install` with `--dependency-update` flag!
