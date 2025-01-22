{{- /*
Copyright (c) Abstract Machines
SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- define "magistrala.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "magistrala.fullname" -}}
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

{{- define "magistrala.gen-certs" -}}
{{- $altNames := list ( printf "%s.%s" (include "magistrala.name" .) .Release.Namespace ) ( printf "%s.%s.svc" (include "magistrala.name" .) .Release.Namespace ) -}}
{{- $ca := genCA "magistrala-ca" 365 -}}
{{- $cert := genSignedCert ( include "magistrala.name" . ) nil $altNames 365 $ca -}}
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
