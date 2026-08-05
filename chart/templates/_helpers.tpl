{{- define "gopad-api.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "gopad-api.fullname" -}}
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

{{- define "gopad-api.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "gopad-api.labels" -}}
helm.sh/chart: "{{ include "gopad-api.chart" . }}"
app.kubernetes.io/name: "gopad-api"
app.kubernetes.io/instance: "{{ .Release.Name }}"
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service | quote }}
{{- with .Values.labels }}
{{ toYaml . }}
{{- end }}
{{- end -}}

{{- define "gopad-api.server.labels" -}}
{{- include "gopad-api.labels" . }}
app.kubernetes.io/component: server
{{- end -}}

{{- define "gopad-api.cleanup.labels" -}}
{{- include "gopad-api.labels" . }}
app.kubernetes.io/component: cleanup
{{- end -}}

{{- define "gopad-api.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default "gopad-api" .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{- define "gopad-api.server.selectorLabels" -}}
app.kubernetes.io/name: gopad-api
app.kubernetes.io/component: server
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "gopad-api.database.secretName" -}}
{{- .Values.config.database.existingSecret | default (printf "%s-database" (include "gopad-api.fullname" .)) -}}
{{- end -}}

{{- define "gopad-api.token.secretName" -}}
{{- .Values.config.token.existingSecret | default (printf "%s-token" (include "gopad-api.fullname" .)) -}}
{{- end -}}

{{- define "gopad-api.admin.secretName" -}}
{{- .Values.config.admin.existingSecret | default (printf "%s-admin" (include "gopad-api.fullname" .)) -}}
{{- end -}}

{{- define "gopad-api.scim.secretName" -}}
{{- .Values.config.scim.existingSecret | default (printf "%s-scim" (include "gopad-api.fullname" .)) -}}
{{- end -}}

{{- define "gopad-api.shared.environment" -}}
- name: GOPAD_API_LOG_LEVEL
  value: "{{ .Values.config.log.level }}"
- name: GOPAD_API_LOG_PRETTY
  value: "false"
- name: GOPAD_API_LOG_COLOR
  value: "false"
{{- if eq .Values.config.database.driver "sqlite3" }}
- name: GOPAD_API_DATABASE_DRIVER
  value: "{{ .Values.config.database.driver }}"
- name: GOPAD_API_DATABASE_NAME
  value: "{{ .Values.config.database.name }}"
{{- else }}
- name: GOPAD_API_DATABASE_DRIVER
  value: "{{ .Values.config.database.driver }}"
- name: GOPAD_API_DATABASE_ADDRESS
  value: "{{ .Values.config.database.address }}"
- name: GOPAD_API_DATABASE_PORT
  value: "{{ .Values.config.database.port }}"
- name: GOPAD_API_DATABASE_NAME
  value: "{{ .Values.config.database.name }}"
- name: GOPAD_API_DATABASE_USERNAME
  value: "{{ .Values.config.database.username }}"
- name: GOPAD_API_DATABASE_PASSWORD
  valueFrom:
    secretKeyRef:
      name: "{{ include "gopad-api.database.secretName" . }}"
      key: "{{ .Values.config.database.passwordKey }}"
{{- end }}
{{- if eq .Values.config.upload.driver "file" }}
- name: GOPAD_API_UPLOAD_DRIVER
  value: "{{ .Values.config.upload.driver }}"
- name: GOPAD_API_UPLOAD_PATH
  value: "{{ .Values.config.upload.path }}"
- name: GOPAD_API_UPLOAD_PERMS
  value: "{{ .Values.config.upload.perms }}"
{{- end }}
{{- if eq .Values.config.upload.driver "s3" }}
- name: GOPAD_API_UPLOAD_DRIVER
  value: "{{ .Values.config.upload.driver }}"
- name: GOPAD_API_UPLOAD_ENDPOINT
  value: "{{ .Values.config.upload.endpoint }}"
- name: GOPAD_API_UPLOAD_BUCKET
  value: "{{ .Values.config.upload.bucket }}"
- name: GOPAD_API_UPLOAD_REGION
  value: "{{ .Values.config.upload.region }}"
- name: GOPAD_API_UPLOAD_PATHSTYLE
  value: "{{ .Values.config.upload.pathstyle }}"
- name: GOPAD_API_UPLOAD_PATH
  value: "{{ .Values.config.upload.path }}"
- name: GOPAD_API_UPLOAD_ACCESS
  valueFrom:
    secretKeyRef:
      name: "{{ include "gopad-api.upload.secretName" . }}"
      key: "{{ .Values.config.upload.accessKey }}"
- name: GOPAD_API_UPLOAD_SECRET
  valueFrom:
    secretKeyRef:
      name: "{{ include "gopad-api.upload.secretName" . }}"
      key: "{{ .Values.config.upload.secretKey }}"
- name: GOPAD_API_UPLOAD_PROXY
  value: "{{ .Values.config.upload.proxy }}"
- name: GOPAD_API_UPLOAD_PERMS
  value: "{{ .Values.config.upload.perms }}"
{{- end }}
{{- end -}}

{{- define "gopad-api.server.environment" -}}
{{- include "gopad-api.shared.environment" . }}
- name: GOPAD_API_SERVER_HOST
  value: "{{ .Values.config.server.host }}"
- name: GOPAD_API_SERVER_ROOT
  value: "{{ .Values.config.server.root }}"
- name: GOPAD_API_SERVER_DOCS
  value: "{{ .Values.config.server.docs }}"
- name: GOPAD_API_TOKEN_EXPIRE
  value: "{{ .Values.config.token.expire }}"
- name: GOPAD_API_TOKEN_SECRET
  valueFrom:
    secretKeyRef:
      name: "{{ include "gopad-api.token.secretName" . }}"
      key: "{{ .Values.config.token.secretKey }}"
- name: GOPAD_API_ADMIN_CREATE
  value: "{{ .Values.config.admin.create }}"
{{- if .Values.config.admin.create }}
- name: GOPAD_API_ADMIN_USERNAME
  valueFrom:
    secretKeyRef:
      name: "{{ include "gopad-api.admin.secretName" . }}"
      key: "{{ .Values.config.admin.usernameKey }}"
- name: GOPAD_API_ADMIN_PASSWORD
  valueFrom:
    secretKeyRef:
      name: "{{ include "gopad-api.admin.secretName" . }}"
      key: "{{ .Values.config.admin.passwordKey }}"
- name: GOPAD_API_ADMIN_EMAIL
  valueFrom:
    secretKeyRef:
      name: "{{ include "gopad-api.admin.secretName" . }}"
      key: "{{ .Values.config.admin.emailKey }}"
{{- end }}
- name: GOPAD_API_SCIM_ENABLED
  value: "{{ .Values.config.scim.enabled }}"
{{- if .Values.config.scim.enabled }}
- name: GOPAD_API_SCIM_TOKEN
  valueFrom:
    secretKeyRef:
      name: "{{ include "gopad-api.scim.secretName" . }}"
      key: "{{ .Values.config.scim.tokenKey }}"
{{- end }}
- name: GOPAD_API_AUTH_CONFIG
  value: "/etc/gopad-api/auth/config.yaml"
{{- end -}}

{{- define "gopad-api.cleanup.environment" -}}
{{- include "gopad-api.shared.environment" . }}
{{- end -}}
