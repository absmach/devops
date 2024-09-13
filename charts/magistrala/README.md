# magistrala

Magistrala IoT Platform

![Version: 0.14.0](https://img.shields.io/badge/Version-0.14.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.14.0](https://img.shields.io/badge/AppVersion-0.14.0-informational?style=flat-square)

**Homepage:** <https://abstractmachines.fr/magistrala.html>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| drasko | <draasko.draskovic@abstractmachines.fr> |  |
| dusan | <dusan.borovcanin@abstractmachines.fr> |  |

## Source Code

* <https://hub.docker.com/u/magistrala>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| @bitnami | postgresqlbootstrap(postgresql) | 12.5.6 |
| @bitnami | postgresqlinvitations(postgresql) | 12.5.6 |
| @bitnami | postgresqlauth(postgresql) | 12.5.6 |
| @bitnami | postgresqlspicedb(postgresql) | 12.5.6 |
| @bitnami | postgresqlthings(postgresql) | 12.5.6 |
| @bitnami | postgresqlusers(postgresql) | 12.5.6 |
| @bitnami | postgresqlui(postgresql) | 12.5.6 |
| @bitnami | postgresqlcerts(postgresql) | 12.5.6 |
| @bitnami | timescaledb(postgresql) | 12.5.6 |
| @bitnami | postgresqljournal(postgresql) | 12.5.6 |
| @bitnami | redis-things(redis) | 19.6.2 |
| @hashicorp | vault(vault) | 0.28.1 |
| @jaegertracing | jaeger | 3.1.1 |
| @nats | nats | 1.2.1 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| adapter_coap.image | object | `{}` |  |
| adapter_coap.port | int | `5683` |  |
| adapter_http.httpPort | int | `8008` |  |
| adapter_http.image | object | `{}` |  |
| auth.accessTokenDuration | string | `"1h"` |  |
| auth.adminEmail | string | `"admin@example.com"` |  |
| auth.adminPassword | string | `"12345678"` |  |
| auth.affinity | object | `{}` |  |
| auth.grpcPort | int | `8181` |  |
| auth.httpPort | int | `8189` |  |
| auth.image | object | `{}` |  |
| auth.invitationDuration | string | `"168h"` |  |
| auth.nodeSelector | object | `{}` |  |
| auth.refreshTokenDuration | string | `"24h"` |  |
| auth.secret | string | `"supersecret"` |  |
| auth.tolerations | object | `{}` |  |
| bootstrap.enabled | bool | `true` |  |
| bootstrap.encKey | string | `"randomstring"` |  |
| bootstrap.eventConsumerName | string | `"EventConsumerByBootstrap"` |  |
| bootstrap.httpPort | int | `9013` |  |
| bootstrap.image | object | `{}` |  |
| bootstrap.redisESPort | int | `6379` |  |
| certs.enabled | bool | `false` |  |
| certs.httpPort | int | `9019` |  |
| certs.image | object | `{}` |  |
| certs.logLevel | string | `"info"` |  |
| certs.signCAKeyPath | string | `"/etc/ssl/certs/ca.key"` |  |
| certs.signCAPath | string | `"/etc/ssl/certs/ca.crt"` |  |
| certs.vault.approleRoleid | string | `""` |  |
| certs.vault.approleSecret | string | `""` |  |
| certs.vault.namespace | string | `""` |  |
| certs.vault.thingsCertsPkiPath | string | `""` |  |
| certs.vault.thingsCertsPkiRoleName | string | `""` |  |
| certs.vault.url | string | `""` |  |
| defaults.eventStreamURL | string | `"magistrala-nats:4222"` |  |
| defaults.image.pullPolicy | string | `"IfNotPresent"` |  |
| defaults.image.rootRepository | string | `"magistrala"` |  |
| defaults.image.tag | string | `"latest"` |  |
| defaults.jaegerCollectorPort | int | `4318` |  |
| defaults.jaegerTraceRatio | int | `10` |  |
| defaults.logLevel | string | `"info"` |  |
| defaults.natsPort | int | `4222` |  |
| defaults.replicaCount | int | `3` |  |
| defaults.sendTelemetry | bool | `true` |  |
| envoy.image.pullPolicy | string | `"IfNotPresent"` |  |
| envoy.image.repository | string | `"envoyproxy/envoy"` |  |
| envoy.image.tag | string | `"v1.31-latest"` |  |
| ingress.annotations | object | `{}` |  |
| ingress.enabled | bool | `true` |  |
| ingress.labels | object | `{}` |  |
| invitations.enabled | bool | `true` |  |
| invitations.httpPort | int | `9020` |  |
| invitations.image | object | `{}` |  |
| jaeger.agent.enabled | bool | `false` |  |
| jaeger.allInOne.enabled | bool | `false` |  |
| jaeger.collector.service.otlp.grpc.name | string | `"otlp-grpc"` |  |
| jaeger.collector.service.otlp.grpc.port | int | `4317` |  |
| jaeger.collector.service.otlp.http.name | string | `"otlp-http"` |  |
| jaeger.collector.service.otlp.http.port | int | `4318` |  |
| jaeger.fullnameOverride | string | `"magistrala-jaeger"` |  |
| jaeger.provisionDataStore.cassandra | bool | `true` |  |
| jaeger.storage.type | string | `"memory"` |  |
| journal.enabled | bool | `true` |  |
| journal.httpPort | int | `9021` |  |
| journal.image | object | `{}` |  |
| mqtt.adapter.image.pullSecrets | object | `{}` |  |
| mqtt.adapter.logLevel | string | `"debug"` |  |
| mqtt.adapter.mqttPort | int | `1884` |  |
| mqtt.adapter.wsPort | int | `8081` |  |
| mqtt.broker.image.repository | string | `"magistrala/vernemq"` |  |
| mqtt.broker.logLevel | string | `"info"` |  |
| mqtt.broker.mqttPort | int | `1883` |  |
| mqtt.broker.persistentVolume.size | string | `"5Gi"` |  |
| mqtt.broker.wsPort | int | `8080` |  |
| mqtt.enabled | bool | `true` |  |
| mqtt.redisCachePort | int | `6379` |  |
| mqtt.redisESPort | int | `6379` |  |
| mqtt.securityContext.fsGroup | int | `10000` |  |
| mqtt.securityContext.runAsGroup | int | `10000` |  |
| mqtt.securityContext.runAsUser | int | `10000` |  |
| nats.config.cluster.enabled | bool | `false` |  |
| nats.config.cluster.replicas | int | `3` |  |
| nats.config.jetstream.enabled | bool | `true` |  |
| nats.config.jetstream.fileStore.enabled | bool | `true` |  |
| nats.config.jetstream.fileStore.pvc.enabled | bool | `true` |  |
| nats.config.jetstream.memoryStore.enabled | bool | `true` |  |
| nats.config.jetstream.memoryStore.maxSize | string | `"2Gi"` |  |
| nginxInternal.image.pullPolicy | string | `"IfNotPresent"` |  |
| nginxInternal.image.repository | string | `"nginx"` |  |
| nginxInternal.image.tag | string | `"1.19.1-alpine"` |  |
| nginxInternal.mtls.intermediateCrt | string | `""` |  |
| nginxInternal.mtls.tls | string | `""` |  |
| postgresqlauth.database | string | `"auth"` |  |
| postgresqlauth.enabled | bool | `true` |  |
| postgresqlauth.global.postgresql.auth.database | string | `"auth"` |  |
| postgresqlauth.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqlauth.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqlauth.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqlauth.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlauth.host | string | `"postgresql-auth"` |  |
| postgresqlauth.name | string | `"postgresql-auth"` |  |
| postgresqlauth.password | string | `"magistrala"` |  |
| postgresqlauth.port | int | `5432` |  |
| postgresqlauth.username | string | `"magistrala"` |  |
| postgresqlbootstrap.database | string | `"bootstrap"` |  |
| postgresqlbootstrap.enabled | bool | `true` |  |
| postgresqlbootstrap.global.postgresql.auth.database | string | `"bootstrap"` |  |
| postgresqlbootstrap.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqlbootstrap.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqlbootstrap.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqlbootstrap.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlbootstrap.host | string | `"postgresql-bootstrap"` |  |
| postgresqlbootstrap.name | string | `"postgresql-bootstrap"` |  |
| postgresqlbootstrap.password | string | `"magistrala"` |  |
| postgresqlbootstrap.port | int | `5432` |  |
| postgresqlbootstrap.username | string | `"magistrala"` |  |
| postgresqlcerts.database | string | `"certs"` |  |
| postgresqlcerts.enabled | bool | `true` |  |
| postgresqlcerts.global.postgresql.auth.database | string | `"certs"` |  |
| postgresqlcerts.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqlcerts.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqlcerts.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqlcerts.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlcerts.host | string | `"postgresql-certs"` |  |
| postgresqlcerts.name | string | `"postgresql-certs"` |  |
| postgresqlcerts.password | string | `"magistrala"` |  |
| postgresqlcerts.port | int | `5432` |  |
| postgresqlcerts.username | string | `"magistrala"` |  |
| postgresqlinvitations.database | string | `"invitations"` |  |
| postgresqlinvitations.enabled | bool | `true` |  |
| postgresqlinvitations.global.postgresql.auth.database | string | `"invitations"` |  |
| postgresqlinvitations.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqlinvitations.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqlinvitations.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqlinvitations.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlinvitations.host | string | `"postgresql-invitations"` |  |
| postgresqlinvitations.name | string | `"postgresql-invitations"` |  |
| postgresqlinvitations.password | string | `"magistrala"` |  |
| postgresqlinvitations.port | int | `5432` |  |
| postgresqlinvitations.username | string | `"magistrala"` |  |
| postgresqljournal.database | string | `"journal"` |  |
| postgresqljournal.enabled | bool | `true` |  |
| postgresqljournal.global.postgresql.auth.database | string | `"journal"` |  |
| postgresqljournal.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqljournal.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqljournal.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqljournal.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqljournal.host | string | `"postgresql-journal"` |  |
| postgresqljournal.name | string | `"postgresql-journal"` |  |
| postgresqljournal.password | string | `"magistrala"` |  |
| postgresqljournal.port | int | `5432` |  |
| postgresqljournal.username | string | `"magistrala"` |  |
| postgresqlspicedb.database | string | `"spicedb"` |  |
| postgresqlspicedb.enabled | bool | `true` |  |
| postgresqlspicedb.global.postgresql.auth.database | string | `"spicedb"` |  |
| postgresqlspicedb.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqlspicedb.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqlspicedb.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqlspicedb.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlspicedb.host | string | `"postgresql-spicedb"` |  |
| postgresqlspicedb.name | string | `"postgresql-spicedb"` |  |
| postgresqlspicedb.password | string | `"magistrala"` |  |
| postgresqlspicedb.port | int | `5432` |  |
| postgresqlspicedb.username | string | `"magistrala"` |  |
| postgresqlthings.database | string | `"things"` |  |
| postgresqlthings.enabled | bool | `true` |  |
| postgresqlthings.global.postgresql.auth.database | string | `"things"` |  |
| postgresqlthings.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqlthings.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqlthings.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqlthings.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlthings.host | string | `"postgresql-things"` |  |
| postgresqlthings.name | string | `"postgresql-things"` |  |
| postgresqlthings.password | string | `"magistrala"` |  |
| postgresqlthings.port | int | `5432` |  |
| postgresqlthings.username | string | `"magistrala"` |  |
| postgresqlui.database | string | `"ui"` |  |
| postgresqlui.enabled | bool | `true` |  |
| postgresqlui.global.postgresql.auth.database | string | `"ui"` |  |
| postgresqlui.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqlui.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqlui.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqlui.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlui.host | string | `"postgresql-ui"` |  |
| postgresqlui.name | string | `"postgresql-ui"` |  |
| postgresqlui.password | string | `"magistrala"` |  |
| postgresqlui.port | int | `5432` |  |
| postgresqlui.username | string | `"magistrala"` |  |
| postgresqlusers.database | string | `"users"` |  |
| postgresqlusers.enabled | bool | `true` |  |
| postgresqlusers.global.postgresql.auth.database | string | `"users"` |  |
| postgresqlusers.global.postgresql.auth.password | string | `"magistrala"` |  |
| postgresqlusers.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| postgresqlusers.global.postgresql.auth.username | string | `"magistrala"` |  |
| postgresqlusers.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlusers.host | string | `"postgresql-users"` |  |
| postgresqlusers.name | string | `"postgresql-users"` |  |
| postgresqlusers.password | string | `"magistrala"` |  |
| postgresqlusers.port | int | `5432` |  |
| postgresqlusers.username | string | `"magistrala"` |  |
| redis-things.cluster.enabled | bool | `false` |  |
| redis-things.usePassword | bool | `false` |  |
| redis-things.volumePermissions.enabled | bool | `true` |  |
| spicedb.affinity | object | `{}` |  |
| spicedb.datastore.engine | string | `"postgres"` |  |
| spicedb.dispatch.enabled | bool | `false` |  |
| spicedb.dispatch.port | int | `50053` |  |
| spicedb.grpc.port | int | `50051` |  |
| spicedb.grpc.presharedKey | string | `"helloworld"` |  |
| spicedb.http.enabled | bool | `false` |  |
| spicedb.http.port | int | `8443` |  |
| spicedb.image.pullSecrets | object | `{}` |  |
| spicedb.image.repository | string | `"authzed/spicedb"` |  |
| spicedb.image.tag | string | `"latest"` |  |
| spicedb.metrics.enabled | bool | `true` |  |
| spicedb.metrics.port | int | `9090` |  |
| spicedb.nodeSelector | object | `{}` |  |
| spicedb.tolerations | object | `{}` |  |
| things.authGrpcPort | int | `7000` |  |
| things.authHttpPort | int | `9001` |  |
| things.httpPort | int | `9000` |  |
| things.image | object | `{}` |  |
| things.redisCachePort | int | `6379` |  |
| things.redisESPort | int | `6379` |  |
| timescaledb.database | string | `"messages"` |  |
| timescaledb.enabled | bool | `true` |  |
| timescaledb.global.postgresql.auth.database | string | `"messages"` |  |
| timescaledb.global.postgresql.auth.password | string | `"magistrala"` |  |
| timescaledb.global.postgresql.auth.postgresPassword | string | `"magistrala"` |  |
| timescaledb.global.postgresql.auth.username | string | `"magistrala"` |  |
| timescaledb.global.postgresql.service.ports.postgresql | int | `5432` |  |
| timescaledb.host | string | `"timescalerw"` |  |
| timescaledb.image.registry | string | `"docker.io"` |  |
| timescaledb.image.repository | string | `"timescale/timescaledb"` |  |
| timescaledb.image.tag | string | `"latest-pg12"` |  |
| timescaledb.name | string | `"timescalerw"` |  |
| timescaledb.password | string | `"magistrala"` |  |
| timescaledb.port | int | `5432` |  |
| timescaledb.reader.enabled | bool | `true` |  |
| timescaledb.reader.http.port | int | `9011` |  |
| timescaledb.reader.image | object | `{}` |  |
| timescaledb.username | string | `"magistrala"` |  |
| timescaledb.writer.enabled | bool | `true` |  |
| timescaledb.writer.http.port | int | `9012` |  |
| timescaledb.writer.image | object | `{}` |  |
| ui.blockKey | string | `"UtgZjr92jwRY6SPUndHXiyl9QY8qTUyZ"` |  |
| ui.contentType | string | `"application/senml+json"` |  |
| ui.enabled | bool | `true` |  |
| ui.googleClientID | string | `""` |  |
| ui.googleClientSecret | string | `""` |  |
| ui.googleRedirectHostname | string | `"https://stage-domain-name"` |  |
| ui.googleRedirectPath | string | `"/oauth/callback/google"` |  |
| ui.googleState | string | `"somerandomstring"` |  |
| ui.hashKey | string | `"5jx4x2Qg9OUmzpP5dbveWQ"` |  |
| ui.image | object | `{}` |  |
| ui.pathPrefix | string | `"/ui"` |  |
| ui.port | int | `9095` |  |
| users.adminEmail | string | `"admin@example.com"` |  |
| users.adminPassword | string | `"12345678"` |  |
| users.allowSelfRegister | bool | `true` |  |
| users.deleteAfter | string | `"720h"` |  |
| users.deleteInterval | string | `"24h"` |  |
| users.grpcPort | int | `7001` |  |
| users.httpPort | int | `9002` |  |
| users.image | object | `{}` |  |
| users.passwordRegex | string | `"^.{8,}$"` |  |
| users.secretKey | string | `"secretKey"` |  |
| users.tokenResetEndpoint | string | `"/reset-request"` |  |
| vault.enabled | bool | `false` |  |
