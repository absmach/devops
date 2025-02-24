{{- /*
Copyright (c) Abstract Machines
SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- define "supermq.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}


{{- define "cassandra.host" -}}
{{- if .Values.provisionDataStore.cassandra -}}
{{- if .Values.storage.cassandra.nameOverride }}
{{- printf "%s" .Values.storage.cassandra.nameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name "cassandra" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- else }}
{{- tpl .Values.storage.cassandra.host . }}
{{- end -}}
{{- end -}}

{{- define "supermq.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "supermq.gen-certs" -}}
{{- $altNames := list ( printf "%s.%s" (include "supermq.name" .) .Release.Namespace ) ( printf "%s.%s.svc" (include "supermq.name" .) .Release.Namespace ) -}}
{{- $ca := genCA "supermq-ca" 365 -}}
{{- $cert := genSignedCert ( include "supermq.name" . ) nil $altNames 365 $ca -}}
tls.crt: {{ $cert.Cert | b64enc }}
tls.key: {{ $cert.Key | b64enc }}
{{- end -}}


{{- define "validateSpiceDBDatastoreEngine" -}}
{{- if and (not (eq . "memory")) (not (eq . "postgres")) -}}
  {{- fail "Invalid value for .Values.spicedb.datastore.engine. Must be 'memory' or 'postgres'." -}}
{{- else -}}
  {{- . -}}
{{- end -}}
{{- end -}}
