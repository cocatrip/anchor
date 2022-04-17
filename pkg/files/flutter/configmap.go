package flutter

// ConfigMap default template
var ConfigMap string = `apiVersion: v1
kind: ConfigMap
metadata:
  annotations:
    service.beta.kubernetes.io/alibaba-cloud-loadbalancer-address-type: "intranet"
  name: {{ include "[[ .Global.APPLICATION_NAME ]].fullname" . }}
  labels:
    {{- include "[[ .Global.APPLICATION_NAME ]].labels" . | nindent 4 }}
data:
  SPRING_ACTIVE_PROFILE: "{{ .Values.config.spring_active_profile}}"`
