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
| nginx_internal.mtls.tls              | TLS secret which contains the server cert/key                              |              |
| nginx_internal.mtls.intermediate_crt | Generic secret which contains the intermediate cert used to verify clients |              |
| ingress.enabled                      | Should the Nginx Ingress be created                                        | true         |
| ingress.hostname                     | Hostname for the Nginx Ingress                                             |              |
| ingress.tls.hostname                 | Hostname of the Nginx Ingress certificate                                  |              |
| ingress.tls.secret                   | TLS secret for the Nginx Ingress                                           |              |
| nats.maxPayload                      | Maximum payload size in bytes that the NATS server will accept             | 268435456    |
| nats.replicaCount                    | NATS replicas                                                              | 3            |
| mqtt.broker.persistentVolume.size    | data Persistent Volume size                                                | 5Gi          |
| influxdb.enabled                     | Enable influxdb reader & writer                                            | false        |
| bootstrap.enabled                    | Enable bootstrap service                                                   | false        |
| adapter_opcua.enabled                | Enable OPC-UA Adapter                                                      | false        |
| twins.enabled                        | Enable twins service                                                       | false        |

All Mainflux services (both core and add-ons) can have their `logLevel`, `image.pullPolicy`, `image.repository` and `image.tag` overridden. 

Mainflux Core is a minimalistic set of required Mainflux services. They are all installed by default:

- adapter_coap
- adapter_http
- adapter_mqtt
- adapter_ws
- things
- ui
- users
- authn

Mainflux Add-ons are optional services that are disabled by default. Find in Configuration table paramaters for enabling them, i.e. to enable influxdb reader & writer you shoud run `helm install` with `--set influxdb=true`.
List of add-ons services in charts:

- bootstrap
- influxdb.writer
- influxdb.reader
- adapter_opcua
- twins

By default scale of MQTT adapter, Things, Envoy, Authn and NATS will be set to 3. It's recommended that you set this values to number of your nodes in Kubernetes cluster, i.e. `--set defaults.replicaCount=3 --set nats.replicaCount=3`

**Note:** make sure you run `helm install` with `--dependency-update` flag!
