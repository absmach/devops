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

| Parameter                              | Description                                                                | Default        |
| -------------------------------------- | -------------------------------------------------------------------------- | -------------- |
| `defaults.logLevel`                    | Log level                                                                  | `debug`        |
| `defaults.image.pullPolicy`            | Docker Image Pull Policy                                                   | `IfNotPresent` |
| `defaults.image.repository`            | Docker Image Repository                                                    | `mainflux`     |
| `defaults.image.tag`                   | Docker Image Tag                                                           | `0.10.0`        |
| `nginx_internal.mtls.tls`              | TLS secret which contains the server cert/key                              | `''`           |
| `nginx_internal.mtls.intermediate_crt` | Generic secret which contains the intermediate cert used to verify clients | `''`           |
| `ingress.enabled`                      | Should the Nginx Ingress be created                                        | `true`         |
| `ingress.hostname`                     | Hostname for the Nginx Ingress                                             | `''`           |
| `ingress.tls.hostname`                 | Hostname of the Nginx Ingress certificate                                  | `''`           |
| `ingress.tls.secret`                   | TLS secret for the Nginx Ingress                                           | `''`           |

All Mainflux services can have their `logLevel`, `image.pullPolicy`, `image.repository` and `image.tag` overridden. The names of the services are:

- adapter_coap
- adapter_http
- adapter_mqtt
- adapter_ws
- things
- ui
- users
- authn
- influxdb_writer
- infuxdb_reader

**Note:** make sure you run `helm install` with `--dependency-update` flag!
