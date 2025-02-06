# Supermq

Event-driven Infrastructure for Modern Cloud

![Version: 0.14.2](https://img.shields.io/badge/Version-0.14.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.14.0](https://img.shields.io/badge/AppVersion-0.14.0-informational?style=flat-square)

**Homepage:** <https://abstractmachines.fr/supermq.html>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| drasko | <drasko.draskovic@abstractmachines.fr> |  |
| dusan | <dusan.borovcanin@abstractmachines.fr> |  |

## Source Code

* <https://hub.docker.com/u/supermq>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| @bitnami | postgresqlauth(postgresql) | 12.5.6 |
| @bitnami | postgresqlspicedb(postgresql) | 12.5.6 |
| @bitnami | postgresqlthings(postgresql) | 12.5.6 |
| @bitnami | postgresqlusers(postgresql) | 12.5.6 |
| @bitnami | postgresqlcerts(postgresql) | 12.5.6 |
| @bitnami | postgresqlinvitations(postgresql) | 12.5.6 |
| @bitnami | postgresqljournal(postgresql) | 12.5.6 |
| @bitnami | postgresqlui(postgresql) | 12.5.6 |
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
| certs.enabled | bool | `true` |  |
| certs.httpPort | int | `9019` |  |
| certs.image | object | `{}` |  |
| certs.logLevel | string | `"info"` |  |
| certs.signCAKeyPath | string | `"/etc/ssl/certs/ca.key"` |  |
| certs.signCAPath | string | `"/etc/ssl/certs/ca.crt"` |  |
| certs.vault.approleRoleid | string | `"supermq"` |  |
| certs.vault.approleSecret | string | `"supermq"` |  |
| certs.vault.namespace | string | `"supermq"` |  |
| certs.vault.thingsCertsPkiPath | string | `"pki_int"` |  |
| certs.vault.thingsCertsPkiRoleName | string | `"supermq_things_certs"` |  |
| certs.vault.url | string | `"http://supermq-vault:8200"` |  |
| clients.authGrpcPort | int | `7000` |  |
| clients.authHttpPort | int | `9001` |  |
| clients.cacheKeyduration | string | `"10m"` |  |
| clients.httpPort | int | `9000` |  |
| clients.image | object | `{}` |  |
| clients.redisCachePort | int | `6379` |  |
| clients.redisESPort | int | `6379` |  |
| defaults.eventStreamURL | string | `"supermq-nats:4222"` |  |
| defaults.image.pullPolicy | string | `"IfNotPresent"` |  |
| defaults.image.rootRepository | string | `"supermq"` |  |
| defaults.image.tag | string | `"latest"` |  |
| defaults.jaegerCollectorPort | int | `4318` |  |
| defaults.jaegerTraceRatio | float | `1` |  |
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
| jaeger.cassandra.persistence.accessModes[0] | string | `"ReadWriteOnce"` |  |
| jaeger.cassandra.persistence.enabled | bool | `true` |  |
| jaeger.cassandra.persistence.size | string | `"10Gi"` |  |
| jaeger.cassandra.persistence.storageClass | string | `"do-block-storage"` |  |
| jaeger.collector.service.otlp.grpc.name | string | `"otlp-grpc"` |  |
| jaeger.collector.service.otlp.grpc.port | int | `4317` |  |
| jaeger.collector.service.otlp.http.name | string | `"otlp-http"` |  |
| jaeger.collector.service.otlp.http.port | int | `4318` |  |
| jaeger.fullnameOverride | string | `"supermq-jaeger"` |  |
| jaeger.provisionDataStore.cassandra | bool | `true` |  |
| jaeger.storage.type | string | `"cassandra"` |  |
| journal.enabled | bool | `true` |  |
| journal.httpPort | int | `9021` |  |
| journal.image | object | `{}` |  |
| mqtt.adapter.image.pullSecrets | object | `{}` |  |
| mqtt.adapter.logLevel | string | `"debug"` |  |
| mqtt.adapter.mqttPort | int | `1884` |  |
| mqtt.adapter.wsPort | int | `8081` |  |
| mqtt.broker.image.repository | string | `"supermq/vernemq"` |  |
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
| postgresqlauth.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlauth.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlauth.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlauth.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlauth.host | string | `"postgresql-auth"` |  |
| postgresqlauth.name | string | `"postgresql-auth"` |  |
| postgresqlauth.password | string | `"supermq"` |  |
| postgresqlauth.port | int | `5432` |  |
| postgresqlauth.username | string | `"supermq"` |  |
| postgresqlcerts.database | string | `"certs"` |  |
| postgresqlcerts.enabled | bool | `true` |  |
| postgresqlcerts.global.postgresql.auth.database | string | `"certs"` |  |
| postgresqlcerts.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlcerts.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlcerts.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlcerts.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlcerts.host | string | `"postgresql-certs"` |  |
| postgresqlcerts.name | string | `"postgresql-certs"` |  |
| postgresqlcerts.password | string | `"supermq"` |  |
| postgresqlcerts.port | int | `5432` |  |
| postgresqlcerts.username | string | `"supermq"` |  |
| postgresqlclients.database | string | `"clients"` |  |
| postgresqlclients.enabled | bool | `true` |  |
| postgresqlclients.global.postgresql.auth.database | string | `"clients"` |  |
| postgresqlclients.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlclients.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlclients.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlclients.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlclients.host | string | `"postgresql-clients"` |  |
| postgresqlclients.name | string | `"postgresql-clients"` |  |
| postgresqlclients.password | string | `"supermq"` |  |
| postgresqlclients.port | int | `5432` |  |
| postgresqlclients.username | string | `"supermq"` |  |
| postgresqlinvitations.database | string | `"invitations"` |  |
| postgresqlinvitations.enabled | bool | `true` |  |
| postgresqlinvitations.global.postgresql.auth.database | string | `"invitations"` |  |
| postgresqlinvitations.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlinvitations.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlinvitations.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlinvitations.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlinvitations.host | string | `"postgresql-invitations"` |  |
| postgresqlinvitations.name | string | `"postgresql-invitations"` |  |
| postgresqlinvitations.password | string | `"supermq"` |  |
| postgresqlinvitations.port | int | `5432` |  |
| postgresqlinvitations.username | string | `"supermq"` |  |
| postgresqljournal.database | string | `"journal"` |  |
| postgresqljournal.enabled | bool | `true` |  |
| postgresqljournal.global.postgresql.auth.database | string | `"journal"` |  |
| postgresqljournal.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqljournal.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqljournal.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqljournal.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqljournal.host | string | `"postgresql-journal"` |  |
| postgresqljournal.name | string | `"postgresql-journal"` |  |
| postgresqljournal.password | string | `"supermq"` |  |
| postgresqljournal.port | int | `5432` |  |
| postgresqljournal.username | string | `"supermq"` |  |
| postgresqlspicedb.database | string | `"spicedb"` |  |
| postgresqlspicedb.enabled | bool | `true` |  |
| postgresqlspicedb.global.postgresql.auth.database | string | `"spicedb"` |  |
| postgresqlspicedb.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlspicedb.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlspicedb.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlspicedb.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlspicedb.host | string | `"postgresql-spicedb"` |  |
| postgresqlspicedb.name | string | `"postgresql-spicedb"` |  |
| postgresqlspicedb.password | string | `"supermq"` |  |
| postgresqlspicedb.port | int | `5432` |  |
| postgresqlspicedb.username | string | `"supermq"` |  |
| postgresqlui.database | string | `"ui"` |  |
| postgresqlui.enabled | bool | `true` |  |
| postgresqlui.global.postgresql.auth.database | string | `"ui"` |  |
| postgresqlui.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlui.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlui.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlui.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlui.host | string | `"postgresql-ui"` |  |
| postgresqlui.name | string | `"postgresql-ui"` |  |
| postgresqlui.password | string | `"supermq"` |  |
| postgresqlui.port | int | `5432` |  |
| postgresqlui.username | string | `"supermq"` |  |
| postgresqlusers.database | string | `"users"` |  |
| postgresqlusers.enabled | bool | `true` |  |
| postgresqlusers.global.postgresql.auth.database | string | `"users"` |  |
| postgresqlusers.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlusers.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlusers.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlusers.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlusers.host | string | `"postgresql-users"` |  |
| postgresqlusers.name | string | `"postgresql-users"` |  |
| postgresqlusers.password | string | `"supermq"` |  |
| postgresqlusers.port | int | `5432` |  |
| postgresqlusers.username | string | `"supermq"` |  |
| redis-clients.cluster.enabled | bool | `false` |  |
| redis-clients.usePassword | bool | `false` |  |
| redis-clients.volumePermissions.enabled | bool | `true` |  |
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
| users.accessTokenDuration | string | `"15m"` |  |
| users.admin.email | string | `"admin@example.com"` |  |
| users.admin.firstname | string | `"super"` |  |
| users.admin.lastname | string | `"admin"` |  |
| users.admin.password | string | `"12345678"` |  |
| users.admin.username | string | `"admin"` |  |
| users.allowSelfRegister | bool | `true` |  |
| users.deleteAfter | string | `"720h"` |  |
| users.deleteInterval | string | `"24h"` |  |
| users.grpcPort | int | `7001` |  |
| users.httpPort | int | `9002` |  |
| users.image | object | `{}` |  |
| users.passwordRegex | string | `"^.{8,}$"` |  |
| users.refreshTokenDuration | string | `"24h"` |  |
| users.secretKey | string | `"HyE2D4RUt9nnKG6v8zKEqAp6g6ka8hhZsqUpzgKvnwpXrNVQSH"` |  |
| users.tokenResetEndpoint | string | `"/reset-request"` |  |
| vault.enabled | bool | `false` |  |
