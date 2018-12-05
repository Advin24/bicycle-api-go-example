{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "fullname" -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Defining base labels to utilize throughout the application.
Can be added to by declaring metadata.labels within Values.yml
*/}}
{{- define "defined.labels" -}}
app: "{{ template "fullname" . }}"
chart: "{{ .Chart.Name }}-{{ .Chart.Version }}"
date: {{ now | htmlDate }}
{{- if .Values.metadata.labels -}}
{{- range $key, $value := .Values.metadata.labels }}
{{ $key }}: {{ $value }}
{{- end }}
{{- end }}
{{- end -}}