# Magistrala Helm Chart

Helm Chart for the Magistrala IoT Platform.

## Prerequisites

- Helm v3
- Stable & Bitnami Helm repos installed
  ```
  helm repo add stable https://charts.helm.sh/stable
  helm repo add bitnami https://charts.bitnami.com/bitnami
  helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
  ```
- Nginx Ingress Controller
- If using the mTLS setup:
  - Certificate/Key installed as**TLS secret**
  - Intermediate certificate installed as**Generic secret**

## Adding the Helm Repository

The Helm charts are published via GitHub Pages. To add the repository to your Helm configuration, run the following command:

```bash
helm repo add devops-charts https://absmach.github.io/devops/
```

Update your local Helm repository cache to fetch the latest charts:

```bash
helm repo update
```

## Installing the Chart

Once the repository is added, you can install the chart using Helm. Replace `<release-name>` with your desired release name:

```bash
helm install <release-name> devops-charts/magistrala
```

This command will install the `magistrala` chart from the GitHub Pages-hosted Helm repository.

## Upgrading the Chart

To upgrade the chart with a new version or updated configuration, use the following command:

```bash
helm upgrade <release-name> devops-charts/magistrala
```

This ensures that your deployment uses the latest version of the chart while retaining any custom configurations.

## Uninstalling the Chart

To uninstall the chart and release, run:

```bash
helm uninstall <release-name>
```

## Configuration

The following table lists the configurable parameters and their default values.

| Parameter                                 | Description                                                                | Default              |
| ----------------------------------------- | -------------------------------------------------------------------------- | -------------------- |
| defaults.logLevel                         | Default log level                                                          | info                 |
| defaults.image.pullPolicy                 | Default Image Pull Policy                                                  | IfNotPresent         |
| defaults.image.rootRepository             | Default Image root repository for magistrala service                       | magistrala           |
| defaults.image.tag                        | Default Image Tag                                                          | latest               |
| defaults.replicaCount                     | Replicas of MQTT adapter, NATS, Things, Envoy, and Auth                    | 3                    |
| defaults.natsPort                         | NATS port                                                                  | 4222                 |
| defaults.jaegerCollectorPort              | Jaeger collector port                                                      | 4318                 |
| defaults.jaegerTraceRatio                 | Jaeger trace ratio (percentage of traces to be sampled)                    | 10                   |
| defaults.sendTelemetry                    | Enable or disable sending telemetry                                        | true                 |
| defaults.eventStreamURL                   | Event stream URL for NATS                                                  | magistrala-nats:4222 |
| ingress.enabled                           | Should the Nginx Ingress be created                                        | true                 |
| ingress.annotations                       | Annotations for the Nginx Ingress                                          | {}                   |
| ingress.labels                            | Labels for the Nginx Ingress                                               | {}                   |
| ingress.hostname                          | Hostname for the Nginx Ingress                                             |                      |
| ingress.tls.hostname                      | TLS hostname for the Nginx Ingress                                         |                      |
| ingress.tls.scret                         | TLS sceret for the Nginx Ingress                                           |                      |
| nginxInternal.image.pullPolicy            | internal nginx image pull policy                                           | IfNotPresent         |
| nginxInternal.image.repository            | inernal nginx image repository                                             | nginx                |
| nginxInternal.image.tag                   | inernal nginx image tag                                                    | 1.19.1-alpine        |
| nginxInternal.mtls.tls                    | TLS secret which contains the server cert/key                              |                      |
| nginxInternal.mtls.intermediateCrt        | Generic secret which contains the intermediate cert used to verify clients |                      |
| envoy.image.pullPolicy                    | Envoy Image Pull Policy for Envoy                                          | IfNotPresent         |
| envoy.image.repository                    | Envoy Image repository for Envoy                                           | envoyproxy/envoy     |
| envoy.image.tag                           | Envoy Image Tag for Envoy                                                  | v1.31-latest         |
| jaeger.fullnameOverride                   | Jaeger fullname override                                                   | magistrala-jaeger    |
| jaeger.provisionDataStore.cassandra       | Provision Jaeger with Cassandra data store                                 | true                 |
| jaeger.agent.enabled                      | Enable Jaeger agent                                                        | false                |
| jaeger.allInOne.enabled                   | Enable Jaeger all-in-one mode                                              | false                |
| jaeger.storage.type                       | Storage type for Jaeger                                                    | memory               |
| jaeger.collector.service.otlp.grpc.port   | OTLP gRPC port for Jaeger collector                                        | 4317                 |
| jaeger.collector.service.otlp.http.port   | OTLP HTTP port for Jaeger collector                                        | 4318                 |
| nats.config.cluster.enabled               | Enable NATS clustering                                                     | false                |
| nats.config.cluster.replicas              | Number of replicas in NATS cluster                                         | 3                    |
| nats.config.jetstream.enabled             | Enable JetStream for NATS                                                  | true                 |
| nats.config.jetstream.fileStore.enabled   | Enable file storage for JetStream                                          | true                 |
| nats.config.jetstream.memoryStore.enabled | Enable memory storage for JetStream                                        | true                 |
| nats.config.jetstream.memoryStore.maxSize | Maximum size for JetStream memory storage                                  | 2Gi                  |
| adapter_coap.image.pullSecrets            | CoAP adapter image pull secrets                                            |                      |
| adapter_coap.image.repository             | CoAP adapter image repository                                              |                      |
| adapter_coap.image.tag                    | CoAP adapter image tag                                                     |                      |
| adapter_coap.image.pullPolicy             | CoAP adapter image pull policy                                             |                      |
| adapter_coap.port                         | CoAP adapter UDP port                                                      | 5683                 |
| adapter_http.image.pullSecrets            | HTTP adapter image pull secrets                                            |                      |
| adapter_http.image.repository             | HTTP adapter image repository                                              |                      |
| adapter_http.image.tag                    | HTTP adapter image tag                                                     |                      |
| adapter_http.image.pullPolicy             | HTTP adapter image pull policy                                             |                      |
| adapter_http.httpPort                     | HTTP adapter port                                                          | 8008                 |
| mqtt.enabled                              | Enable MQTT adapter                                                        | true                 |
| mqtt.securityContext.runAsUser            | Run MQTT adapter as a specific user                                        | 10000                |
| mqtt.securityContext.runAsGroup           | Run MQTT adapter as a specific group                                       | 10000                |
| mqtt.securityContext.fsGroup              | Filesystem group for MQTT adapter                                          | 10000                |
| mqtt.adapter.image.pullSecrets            | MQTT adapter image pull secrets                                            |                      |
| mqtt.adapter.image.repository             | MQTT adapter image repository                                              |                      |
| mqtt.adapter.image.tag                    | MQTT adapter image tag                                                     |                      |
| mqtt.adapter.image.pullPolicy             | MQTT adapter image pull policy                                             |                      |
| mqtt.adapter.mqttPort                     | MQTT adapter MQTT port                                                     | 1884                 |
| mqtt.adapter.wsPort                       | MQTT adapter WebSocket port                                                | 8081                 |
| mqtt.adapter.logLevel                     | Log level for MQTT adapter                                                 | debug                |
| mqtt.broker.mqttPort                      | MQTT broker MQTT port                                                      | 1883                 |
| mqtt.broker.wsPort                        | MQTT broker WebSocket port                                                 | 8080                 |
| mqtt.broker.logLevel                      | Log level for MQTT broker                                                  | info                 |
| mqtt.broker.persistentVolume.size         | Persistent volume size for MQTT broker                                     | 5Gi                  |
| mqtt.redisESPort                          | MQTT adapter Redis Event Store port                                        | 6379                 |
| mqtt.redisCachePort                       | MQTT adapter Redis Cache port                                              | 6379                 |
| spicedb.image.repository                  | Docker Image repository for SpiceDB                                        | authzed/spicedb      |
| spicedb.image.tag                         | Docker Image Tag for SpiceDB                                               | latest               |
| spicedb.grpc.presharedKey                 | Pre-shared key for SpiceDB gRPC                                            | helloworld           |
| spicedb.grpc.port                         | SpiceDB gRPC port                                                          | 50051                |
| spicedb.datastore.engine                  | Datastore engine for SpiceDB                                               | postgres             |
| spicedb.dispatch.port                     | Dispatch port for SpiceDB                                                  | 50053                |
| spicedb.dispatch.enabled                  | Enable Dispatch for SpiceDB                                                | false                |
| spicedb.http.enabled                      | Enable HTTP for SpiceDB                                                    | false                |
| spicedb.http.port                         | HTTP port for SpiceDB                                                      | 8443                 |
| spicedb.metrics.enabled                   | Enable metrics for SpiceDB                                                 | true                 |
| spicedb.metrics.port                      | Metrics port for SpiceDB                                                   | 9090                 |
| postgresqlspicedb.enabled                 | Enable PostgreSQL for SpiceDB                                              | true                 |
| postgresqlspicedb.host                    | Host for PostgreSQL SpiceDB                                                | postgresql-spicedb   |
| postgresqlspicedb.port                    | PostgreSQL port for SpiceDB                                                | 5432                 |
| postgresqlspicedb.database                | Database name for SpiceDB                                                  | spicedb              |
| postgresqlspicedb.username                | Username for PostgreSQL SpiceDB                                            | magistrala           |
| postgresqlspicedb.password                | Password for PostgreSQL SpiceDB                                            | magistrala           |
| auth.httpPort                             | HTTP port for Auth service                                                 | 8189                 |
| auth.grpcPort                             | gRPC port for Auth service                                                 | 8181                 |
| auth.secret                               | Secret key for Auth service                                                | supersecret          |
| auth.adminEmail                           | Admin email for Auth service                                               | admin@example.com    |
| auth.adminPassword                        | Admin password for Auth service                                            | 12345678             |
| auth.accessTokenDuration                  | Access token duration for Auth service                                     | 1h                   |
| auth.refreshTokenDuration                 | Refresh token duration for Auth service                                    | 24h                  |
| auth.invitationDuration                   | Invitation duration for Auth service                                       | 168h                 |
| postgresqlauth.enabled                    | Enable PostgreSQL for Auth service                                         | true                 |
| postgresqlauth.host                       | Host for PostgreSQL Auth service                                           | postgresql-auth      |
| postgresqlauth.port                       | PostgreSQL port for Auth service                                           | 5432                 |
| postgresqlauth.database                   | Database name for Auth service                                             | auth                 |
| postgresqlauth.username                   | Username for PostgreSQL Auth service                                       | magistrala           |
| postgresqlauth.password                   | Password for PostgreSQL Auth service                                       | magistrala           |
| users.dbPort                              | Users service DB port                                                      | 5432                 |
| users.httpPort                            | Users service HTTP port                                                    | 9002                 |
| users.grpcPort                            | Users service gRPC port                                                    | 7001                 |
| things.dbPort                             | Things service DB port                                                     | 5432                 |
| things.httpPort                           | Things service HTTP port                                                   | 9000                 |
| things.authGrpcPort                       | Things service Auth gRPC port                                              | 7000                 |
| things.authHttpPort                       | Things service Auth HTTP port                                              | 9001                 |
| things.redisESPort                        | Things service Redis Event Store port                                      | 6379                 |
| things.redisCachePort                     | Things service Redis Auth Cache port                                       | 6379                 |
| adapter_http.httpPort                     | HTTP adapter port                                                          | 8008                 |
| mqtt.adapter.mqttPort                     | MQTT adapter port                                                          | 1884                 |
| mqtt.adapter.wsPort                       | MQTT adapter WS port                                                       | 8081                 |
| mqtt.broker.mqttPort                      | MQTT adapter broker port                                                   | 1883                 |
| mqtt.broker.wsPort                        | MQTT adapter broker WS port                                                | 8080                 |
| mqtt.broker.persistentVolume.size         | MQTT adapter broker data Persistent Volume size                            | 5Gi                  |
| mqtt.redisESPort                          | MQTT adapter Event Store port                                              | 6379                 |
| mqtt.redisCachePort                       | MQTT adapter Redis Auth Cache port                                         | 6379                 |
| adapter_coap.udpPort                      | CoAP adapter UDP port                                                      | 5683                 |
| ui.port                                   | UI port                                                                    | 3000                 |
| bootstrap.enabled                         | Enable bootstrap service                                                   | false                |
| bootstrap.dbPort                          | Bootstrap service DB port                                                  | 5432                 |
| bootstrap.httpPort                        | Bootstrap service HTTP port                                                | 9013                 |
| bootstrap.redisESPort                     | Bootstrap service Redis Event Store port                                   | 6379                 |
| influxdb.enabled                          | Enable InfluxDB reader & writer                                            | false                |
| influxdb.dbPort                           | InfluxDB port                                                              | 8086                 |
| influxdb.writer.httpPort                  | InfluxDB writer HTTP port                                                  | 9006                 |
| influxdb.reader.httpPort                  | InfluxDB reader HTTP port                                                  | 9005                 |
| influxdb.backup.enabled                   | Enable InfluxDB backup                                                     | false                |
| influxdb.backup.cronjob.schedule          | Crontab style time schedule for backup execution                           | "0 2 \* \* \*"       |
| adapter_opcua.enabled                     | Enable OPC-UA adapter                                                      | false                |
| adapter_opcua.httpPort                    | OPC-UA adapter HTTP port                                                   | 8188                 |
| adapter_opcua.redisRouteMapPort           | OPC-UA adapter Redis Auth Cache port                                       | 6379                 |
| adapter_lora.enabled                      | Enable LoRa adapter                                                        | false                |
| adapter_lora.httpPort                     | LoRa adapter HTTP port                                                     | 9017                 |
| adapter_lora.redisRouteMapPort            | LoRa adapter Redis Auth Cache port                                         | 6379                 |
| twins.enabled                             | Enable twins service                                                       | false                |
| twins.dbPort                              | Twins service DB port                                                      | 27017                |
| twins.httpPort                            | Twins service HTTP port                                                    | 9018                 |
| twins.redisCachePort                      | Twins service Redis Cache port                                             | 6379                 |
| certs.enabled                             | Enable certs service                                                       | false                |
| notifier_smtp.enabled                     | Enable SMTP notifier                                                       | false                |
| notifier_smtp.emailHost                   | SMTP host                                                                  | false                |
| notifier_smtp.smtpPort                    | SMTP port                                                                  | false                |
| notifier_smtp.fromName                    | SMTP notifier `from` name                                                  | false                |
| notifier_smtp.fromEmail                   | SMTP `from` email address                                                  | false                |
| notifier_smtp.username                    | SMTP username                                                              | false                |
| notifier_smtp.password                    | SMTP password                                                              | false                |
| notifier_smtp.secret                      | SMTP secret                                                                | false                |
| notifier_smtp.httpPort                    | SMTP notifier HTTP port                                                    | false                |
| loki_stack.enabled                        | Enable Loki_Stack                                                          | true                 |

All Magistrala services (both core and add-ons) can have their `logLevel`, `image.pullPolicy`, `image.repository` and `image.tag` overridden.

Magistrala Core is a minimalistic set of required Magistrala services. They are all installed by default:

- auth
- users
- things
- adapter_http
- adapter_mqtt
- adapter_coap
- ui

Magistrala Add-ons are optional services that are disabled by default. Find in Configuration table parameters for enabling them, i.e. to enable influxdb reader & writer you should run `helm install` with `--set influxdb=true`.
List of add-ons services in charts:

- bootstrap
- influxdb.writer
- influxdb.reader
- adapter_opcua
- adapter_lora
- twins
- notifier_smtp

By default scale of MQTT adapter, Things, Envoy, Auth and NATS will be set to 3. It's recommended that you set this values to number of your nodes in Kubernetes cluster, i.e. `--set defaults.replicaCount=3 --set nats.replicaCount=3`

**Note:** make sure you run `helm install` with `--dependency-update` flag!
