package files

// Service default template
var Service string = `apiVersion: v1
kind: Service
metadata:
  annotations:
    service.beta.kubernetes.io/alibaba-cloud-loadbalancer-address-type: "intranet"
    #service.beta.kubernetes.io/alibaba-cloud-loadbalancer-vswitch-id: "vsw-k1ajdhoii6qauydjn44qv"
  name: {{ include "[[ .Global.APPLICATION_NAME ]].fullname" . }}
  labels:
    {{- include "[[ .Global.APPLICATION_NAME ]].labels" . | nindent 4 }}
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: {{ .Values.service.targetport }}
      protocol: TCP
      name: http
  selector:
    {{- include "[[ .Global.APPLICATION_NAME ]].selectorLabels" . | nindent 4 }}`