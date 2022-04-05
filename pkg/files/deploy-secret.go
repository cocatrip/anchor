package files

var deploymentSecret string = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "ad1-adm-main.fullname" . }}
  labels:
    {{- include "ad1-adm-main.labels" . | nindent 4 }}
spec:
  {{- if not .Values.autoscaling.enabled }}
  replicas: {{ .Values.replicaCount }}
  {{- end }}
  selector:
    matchLabels:
      {{- include "ad1-adm-main.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      {{- with .Values.podAnnotations }}
      annotations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      labels:
        {{- include "ad1-adm-main.selectorLabels" . | nindent 8 }}
    spec:
      {{- with .Values.imagePullSecrets }}
      imagePullSecrets:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      serviceAccountName: {{ include "ad1-adm-main.serviceAccountName" . }}
      securityContext:
        {{- toYaml .Values.podSecurityContext | nindent 8 }}
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.securityContext | nindent 12 }}
            #image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          env:
            - name: SPRING_ACTIVE_PROFILE
              valueFrom:
                  configMapKeyRef:
                      key: SPRING_ACTIVE_PROFILE
                      name: {{ .Release.Name }}
            - name: DB_USER
              valueFrom:
                  secretKeyRef:
                      key: DB_USER
                      name: {{ .Values.config.secret_name }}
            - name: DB_PASSWORD
              valueFrom:
                  secretKeyRef:
                      key: DB_PASSWORD
                      name: {{ .Values.config.secret_name }}
          ports:
            - name: http
              containerPort: {{ .Values.service.targetport }}
              protocol: TCP
          #livenessProbe:
           # httpGet:
            #  path: /
             # port: http
          #readinessProbe:
           # httpGet:
            #  path: /
             # port: http
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
      {{- end }}`
