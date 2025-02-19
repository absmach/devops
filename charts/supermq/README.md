# Supermq

Event-driven Infrastructure for Modern Cloud

![Version: 0.16.0](https://img.shields.io/badge/Version-0.16.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.16.0](https://img.shields.io/badge/AppVersion-0.16.0-informational?style=flat-square)

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
| @bitnami | postgresqlusers(postgresql) | 12.5.6 |
| @bitnami | postgresqlchannels(postgresql) | 12.5.6 |
| @bitnami | postgresqlui(postgresql) | 12.5.6 |
| @bitnami | postgresqlspicedb(postgresql) | 12.5.6 |
| @bitnami | postgresqlcerts(postgresql) | 12.5.6 |
| @bitnami | postgresqlclients(postgresql) | 12.5.6 |
| @bitnami | postgresqldomains(postgresql) | 12.5.6 |
| @bitnami | postgresqljournal(postgresql) | 12.5.6 |
| @bitnami | postgresqlauth(postgresql) | 12.5.6 |
| @bitnami | postgresqlgroups(postgresql) | 12.5.6 |
| @bitnami | redis-clients(redis) | 19.6.2 |
| @hashicorp | vault(vault) | 0.28.1 |
| @jaegertracing | jaeger | 3.1.1 |
| @nats | nats | 1.2.1 |
| https://fluent.github.io/helm-charts | fluent-bit(fluent-bit) | 0.48.5 |
| https://grafana.github.io/helm-charts | grafana(grafana) | 8.9.0 |
| https://prometheus-community.github.io/helm-charts | prometheus(prometheus) | 27.3.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| adapter_coap.image | object | `{}` |  |
| adapter_coap.jaegerTraceRatio | float | `1` |  |
| adapter_coap.port | int | `5683` |  |
| adapter_coap.sendTelemetry | bool | `true` |  |
| adapter_http.httpPort | int | `8008` |  |
| adapter_http.image | object | `{}` |  |
| adapter_ws.httpPort | int | `8186` |  |
| adapter_ws.image | object | `{}` |  |
| auth.accessTokenDuration | string | `"1h"` |  |
| auth.adminEmail | string | `"admin@example.com"` |  |
| auth.adminPassword | string | `"12345678"` |  |
| auth.affinity | object | `{}` |  |
| auth.grpcClientCert | string | `"./ssl/certs/auth-grpc-client.crt"` |  |
| auth.grpcClientKey | string | `"./ssl/certs/auth-grpc-client.key"` |  |
| auth.grpcPort | int | `7001` |  |
| auth.grpcTimeout | string | `"300s"` |  |
| auth.httpPort | int | `9001` |  |
| auth.image | object | `{}` |  |
| auth.nodeSelector | object | `{}` |  |
| auth.refreshTokenDuration | string | `"24h"` |  |
| auth.secretKey | string | `"supersecret"` |  |
| auth.tolerations | object | `{}` |  |
| certs.enabled | bool | `true` |  |
| certs.httpPort | int | `9019` |  |
| certs.image | object | `{}` |  |
| certs.logLevel | string | `"error"` |  |
| certs.sdkCertsUrl | string | `"${SMQ_CERTS_SDK_HOST}:9010"` |  |
| certs.sdkHost | string | `"http://supermq-am-certs"` |  |
| certs.sdkTlsVerification | string | `"false"` |  |
| certs.signCAKeyPath | string | `"/etc/ssl/certs/ca.key"` |  |
| certs.signCAPath | string | `"/etc/ssl/certs/ca.crt"` |  |
| certs.vault.approleRoleid | string | `"supermq"` |  |
| certs.vault.approleSecret | string | `"supermq"` |  |
| certs.vault.namespace | string | `"supermq"` |  |
| certs.vault.thingsCertsPkiPath | string | `"pki_int"` |  |
| certs.vault.thingsCertsPkiRoleName | string | `"supermq_things_certs"` |  |
| certs.vault.url | string | `"http://supermq-vault:8200"` |  |
| channels.grpcClientCaCerts | string | `"./ssl/certs/ca.crt"` |  |
| channels.grpcClientCert | string | `"./ssl/certs/channels-grpc-client.crt"` |  |
| channels.grpcClientKey | string | `"./ssl/certs/channels-grpc-client.key"` |  |
| channels.grpcPort | int | `7005` |  |
| channels.grpcServerCert | string | `"./ssl/certs/channels-grpc-server.crt"` |  |
| channels.grpcServerKey | string | `"./ssl/certs/channels-grpc-server.key"` |  |
| channels.grpcTimeout | string | `"1s"` |  |
| channels.httpPort | int | `9005` |  |
| channels.image | object | `{}` |  |
| clients.authGrpcPort | int | `7006` |  |
| clients.authHttpPort | int | `9001` |  |
| clients.cacheKeyduration | string | `"10m"` |  |
| clients.grpcClientCert | string | `"./ssl/certs/clients-grpc-client.crt"` |  |
| clients.grpcClientKey | string | `"./ssl/certs/clients-grpc-client.key"` |  |
| clients.grpcTimeout | string | `"1s"` |  |
| clients.httpPort | int | `9006` |  |
| clients.image | object | `{}` |  |
| clients.redisCachePort | int | `6379` |  |
| clients.redisESPort | int | `6379` |  |
| defaults.eventStreamURL | string | `"supermq-nats:4222"` |  |
| defaults.image.pullPolicy | string | `"IfNotPresent"` |  |
| defaults.image.rootRepository | string | `"supermq"` |  |
| defaults.image.tag | string | `"latest"` |  |
| defaults.jaegerCollectorPort | int | `4318` |  |
| defaults.jaegerTraceRatio | float | `1` |  |
| defaults.logLevel | string | `"error"` |  |
| defaults.natsPort | int | `4222` |  |
| defaults.replicaCount | int | `3` |  |
| defaults.sendTelemetry | bool | `true` |  |
| domains.cacheKeyduration | string | `"10m"` |  |
| domains.grpcClientCaCerts | string | `"./ssl/certs/ca.crt"` |  |
| domains.grpcClientCert | string | `"./ssl/certs/domains-grpc-client.crt"` |  |
| domains.grpcPort | int | `7003` |  |
| domains.grpcTimeout | string | `"300s"` |  |
| domains.httpPort | int | `9003` |  |
| domains.image | object | `{}` |  |
| domains.redisTCPPort | int | `6379` |  |
| envoy.image.pullPolicy | string | `"IfNotPresent"` |  |
| envoy.image.repository | string | `"envoyproxy/envoy"` |  |
| envoy.image.tag | string | `"v1.31-latest"` |  |
| fluent-bit.config.filters | string | `"[FILTER]\n    Name         kubernetes\n    Match        kube.*\n    k8s-logging.exclude off\n    Buffer_Size 256k\n"` |  |
| fluent-bit.config.inputs | string | `"[INPUT]\n    Name             tail\n    Path             /var/log/containers/*.log\n    Read_from_head   true\n    Tag              kube.*\n"` |  |
| fluent-bit.config.outputs | string | `"[OUTPUT]\n    Name        loki\n    Match       *\n    Host        supermq-loki.loki\n    Port        3100\n    Uri         /loki/api/v1/push\n    Labels      job=fluent-bit\n    Label_Keys  $kubernetes['namespace_name'], $kubernetes['pod_name']\n    Line_Format json\n    Auto_Kubernetes_Labels off\n"` |  |
| fluent-bit.enabled | bool | `true` |  |
| fluent-bit.resources | object | `{}` |  |
| fluent-bit.serviceAccount.create | bool | `true` |  |
| grafana.adminPassword | string | `"12345678"` |  |
| grafana.adminUser | string | `"admin"` |  |
| grafana.datasources."datasources.yaml".apiVersion | int | `1` |  |
| grafana.datasources."datasources.yaml".datasources[0].access | string | `"proxy"` |  |
| grafana.datasources."datasources.yaml".datasources[0].isDefault | bool | `true` |  |
| grafana.datasources."datasources.yaml".datasources[0].name | string | `"Prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[0].type | string | `"prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[0].url | string | `"http://supermq-prometheus-server:9200"` |  |
| grafana.datasources."datasources.yaml".datasources[1].access | string | `"proxy"` |  |
| grafana.datasources."datasources.yaml".datasources[1].isDefault | bool | `false` |  |
| grafana.datasources."datasources.yaml".datasources[1].name | string | `"Loki"` |  |
| grafana.datasources."datasources.yaml".datasources[1].type | string | `"loki"` |  |
| grafana.datasources."datasources.yaml".datasources[1].url | string | `"http://supermq-loki.loki:3100"` |  |
| grafana.enabled | bool | `true` |  |
| grafana.service.type | string | `"LoadBalancer"` |  |
| groups.grpcClientCaCerts | string | `"./ssl/certs/ca.crt"` |  |
| groups.grpcClientCert | string | `"./ssl/certs/groups-grpc-client.crt"` |  |
| groups.grpcClientKey | string | `"./ssl/certs/groups-grpc-client.key"` |  |
| groups.grpcPort | int | `7004` |  |
| groups.grpcServerCert | string | `"./ssl/certs/groups-grpc-server.crt"` |  |
| groups.grpcServerKey | string | `"./ssl/certs/groups-grpc-server.key"` |  |
| groups.grpcTimeout | string | `"300s"` |  |
| groups.httpPort | int | `9004` |  |
| groups.image | object | `{}` |  |
| ingress.annotations | object | `{}` |  |
| ingress.enabled | bool | `true` |  |
| ingress.labels | object | `{}` |  |
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
| mqtt.adapter.forwarderTimeout | string | `"30s"` |  |
| mqtt.adapter.image.pullSecrets | object | `{}` |  |
| mqtt.adapter.logLevel | string | `"error"` |  |
| mqtt.adapter.mqttPort | int | `1884` |  |
| mqtt.adapter.qos | string | `"2"` |  |
| mqtt.adapter.wsPort | int | `8081` |  |
| mqtt.broker.image.repository | string | `"supermq/vernemq"` |  |
| mqtt.broker.logLevel | string | `"error"` |  |
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
| postgresqlchannels.database | string | `"channels"` |  |
| postgresqlchannels.enabled | bool | `true` |  |
| postgresqlchannels.global.postgresql.auth.database | string | `"channels"` |  |
| postgresqlchannels.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlchannels.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlchannels.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlchannels.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlchannels.host | string | `"channels-db"` |  |
| postgresqlchannels.name | string | `"postgresql-channels"` |  |
| postgresqlchannels.password | string | `"supermq"` |  |
| postgresqlchannels.port | int | `5432` |  |
| postgresqlchannels.username | string | `"supermq"` |  |
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
| postgresqldomains.database | string | `"domains"` |  |
| postgresqldomains.enabled | bool | `true` |  |
| postgresqldomains.global.postgresql.auth.database | string | `"domains"` |  |
| postgresqldomains.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqldomains.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqldomains.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqldomains.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqldomains.host | string | `"postgresql-domains"` |  |
| postgresqldomains.name | string | `"postgresql-domains"` |  |
| postgresqldomains.password | string | `"supermq"` |  |
| postgresqldomains.port | int | `5432` |  |
| postgresqldomains.username | string | `"supermq"` |  |
| postgresqlgroups.database | string | `"groups"` |  |
| postgresqlgroups.enabled | bool | `true` |  |
| postgresqlgroups.global.postgresql.auth.database | string | `"groups"` |  |
| postgresqlgroups.global.postgresql.auth.password | string | `"supermq"` |  |
| postgresqlgroups.global.postgresql.auth.postgresPassword | string | `"supermq"` |  |
| postgresqlgroups.global.postgresql.auth.username | string | `"supermq"` |  |
| postgresqlgroups.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresqlgroups.host | string | `"postgresql-groups"` |  |
| postgresqlgroups.name | string | `"postgresql-groups"` |  |
| postgresqlgroups.password | string | `"supermq"` |  |
| postgresqlgroups.port | int | `5432` |  |
| postgresqlgroups.username | string | `"supermq"` |  |
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
| prometheus.alertmanager.enabled | bool | `true` |  |
| prometheus.alertmanager.persistence.size | string | `"2Gi"` |  |
| prometheus.configmapReload.prometheus.containerPort | int | `8080` |  |
| prometheus.configmapReload.prometheus.containerPortName | string | `"metrics"` |  |
| prometheus.configmapReload.prometheus.enabled | bool | `true` |  |
| prometheus.configmapReload.prometheus.image.pullPolicy | string | `"IfNotPresent"` |  |
| prometheus.configmapReload.prometheus.image.repository | string | `"quay.io/prometheus-operator/prometheus-config-reloader"` |  |
| prometheus.configmapReload.prometheus.image.tag | string | `"v0.79.2"` |  |
| prometheus.configmapReload.prometheus.name | string | `"configmap-reload"` |  |
| prometheus.configmapReload.prometheus.resources | object | `{}` |  |
| prometheus.enabled | bool | `true` |  |
| prometheus.kubeStateMetrics.enabled | bool | `true` |  |
| prometheus.nodeExporter.containerPort | int | `9100` |  |
| prometheus.nodeExporter.enabled | bool | `true` |  |
| prometheus.nodeExporter.extraArgs[0] | string | `"--web.listen-address=0.0.0.0:9100"` |  |
| prometheus.nodeExporter.hostNetwork | bool | `true` |  |
| prometheus.nodeExporter.hostPID | bool | `true` |  |
| prometheus.nodeExporter.hostPort | int | `9100` |  |
| prometheus.nodeExporter.image.pullPolicy | string | `"IfNotPresent"` |  |
| prometheus.nodeExporter.image.repository | string | `"quay.io/prometheus/node-exporter"` |  |
| prometheus.nodeExporter.image.tag | string | `"v1.8.2"` |  |
| prometheus.nodeExporter.livenessProbe.httpGet.path | string | `"/metrics"` |  |
| prometheus.nodeExporter.livenessProbe.httpGet.port | int | `9100` |  |
| prometheus.nodeExporter.livenessProbe.httpGet.scheme | string | `"HTTP"` |  |
| prometheus.nodeExporter.livenessProbe.initialDelaySeconds | int | `5` |  |
| prometheus.nodeExporter.livenessProbe.periodSeconds | int | `10` |  |
| prometheus.nodeExporter.nodeSelector."kubernetes.io/os" | string | `"linux"` |  |
| prometheus.nodeExporter.readinessProbe.httpGet.path | string | `"/metrics"` |  |
| prometheus.nodeExporter.readinessProbe.httpGet.port | int | `9100` |  |
| prometheus.nodeExporter.readinessProbe.httpGet.scheme | string | `"HTTP"` |  |
| prometheus.nodeExporter.readinessProbe.initialDelaySeconds | int | `5` |  |
| prometheus.nodeExporter.readinessProbe.periodSeconds | int | `10` |  |
| prometheus.nodeExporter.service.annotations."prometheus.io/scrape" | string | `"true"` |  |
| prometheus.nodeExporter.service.clusterIP | string | `""` |  |
| prometheus.nodeExporter.service.enabled | bool | `true` |  |
| prometheus.nodeExporter.service.port | int | `9100` |  |
| prometheus.nodeExporter.service.servicePort | int | `9100` |  |
| prometheus.nodeExporter.service.targetPort | int | `9100` |  |
| prometheus.nodeExporter.service.type | string | `"ClusterIP"` |  |
| prometheus.nodeExporter.tolerations[0].effect | string | `"NoSchedule"` |  |
| prometheus.nodeExporter.tolerations[0].key | string | `"node-role.kubernetes.io/master"` |  |
| prometheus.nodeExporter.tolerations[0].operator | string | `"Exists"` |  |
| prometheus.prometheusPushgateway.enabled | bool | `false` |  |
| prometheus.pushgateway.enabled | bool | `false` |  |
| prometheus.rbac.create | bool | `true` |  |
| prometheus.server.extraFlags[0] | string | `"web.enable-lifecycle"` |  |
| prometheus.server.image.pullPolicy | string | `"IfNotPresent"` |  |
| prometheus.server.image.repository | string | `"quay.io/prometheus/prometheus"` |  |
| prometheus.server.image.tag | string | `""` |  |
| prometheus.server.ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| prometheus.server.ingress.enabled | bool | `true` |  |
| prometheus.server.ingress.hosts[0] | string | `"prometheus.example.com"` |  |
| prometheus.server.ingress.ingressClassName | string | `"nginx"` |  |
| prometheus.server.livenessProbe.httpGet.path | string | `"/-/healthy"` |  |
| prometheus.server.livenessProbe.httpGet.port | int | `9090` |  |
| prometheus.server.livenessProbe.httpGet.scheme | string | `"HTTP"` |  |
| prometheus.server.livenessProbe.initialDelaySeconds | int | `30` |  |
| prometheus.server.livenessProbe.timeoutSeconds | int | `5` |  |
| prometheus.server.name | string | `"server"` |  |
| prometheus.server.persistentVolume.accessModes[0] | string | `"ReadWriteOnce"` |  |
| prometheus.server.persistentVolume.enabled | bool | `true` |  |
| prometheus.server.persistentVolume.mountPath | string | `"/data"` |  |
| prometheus.server.persistentVolume.size | string | `"8Gi"` |  |
| prometheus.server.persistentVolume.storageClass | string | `"do-block-storage"` |  |
| prometheus.server.readinessProbe.httpGet.path | string | `"/-/ready"` |  |
| prometheus.server.readinessProbe.httpGet.port | int | `9090` |  |
| prometheus.server.readinessProbe.httpGet.scheme | string | `"HTTP"` |  |
| prometheus.server.readinessProbe.initialDelaySeconds | int | `5` |  |
| prometheus.server.readinessProbe.timeoutSeconds | int | `5` |  |
| prometheus.server.resources | object | `{}` |  |
| prometheus.server.securityContext.fsGroup | int | `65534` |  |
| prometheus.server.securityContext.runAsGroup | int | `65534` |  |
| prometheus.server.securityContext.runAsNonRoot | bool | `true` |  |
| prometheus.server.securityContext.runAsUser | int | `65534` |  |
| prometheus.server.service.annotations."prometheus.io/scrape" | string | `"true"` |  |
| prometheus.server.service.clusterIP | string | `""` |  |
| prometheus.server.service.enabled | bool | `true` |  |
| prometheus.server.service.port | int | `9200` |  |
| prometheus.server.service.portName | string | `"metrics"` |  |
| prometheus.server.service.servicePort | int | `9200` |  |
| prometheus.server.service.targetPort | int | `9090` |  |
| prometheus.server.service.type | string | `"ClusterIP"` |  |
| prometheus.serviceAccounts.server.annotations | object | `{}` |  |
| prometheus.serviceAccounts.server.automountServiceAccountToken | bool | `true` |  |
| prometheus.serviceAccounts.server.create | bool | `true` |  |
| prometheus.serviceAccounts.server.name | string | `""` |  |
| redis-clients.cluster.enabled | bool | `false` |  |
| redis-clients.usePassword | bool | `false` |  |
| redis-clients.volumePermissions.enabled | bool | `true` |  |
| spicedb.affinity | object | `{}` |  |
| spicedb.datastore.engine | string | `"postgres"` |  |
| spicedb.dispatch.enabled | bool | `false` |  |
| spicedb.dispatch.port | int | `50053` |  |
| spicedb.grpc.port | int | `50051` |  |
| spicedb.grpc.presharedKey | string | `"12345678"` |  |
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
| users.httpPort | int | `9002` |  |
| users.image | object | `{}` |  |
| users.passwordRegex | string | `"^.{8,}$"` |  |
| users.refreshTokenDuration | string | `"24h"` |  |
| users.secretKey | string | `"supersecret"` |  |
| users.tokenResetEndpoint | string | `"/reset-request"` |  |
| vault.enabled | bool | `false` |  |
