package node

var Secret string = `apiVersion: v1
kind: Secret
metadata:
# Secret name for kubernetes secret
  name: {{ .Values.config.secret_name }}
type: Opaque
data:
  DB_USER: {{ .Values.config.DB_USER }}
  DB_PASSWORD: {{ .Values.config.DB_PASSWORD }}`
