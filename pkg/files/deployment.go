package files

var Deployment string = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "[[ .Global.APPLICATION_NAME ]].fullname" . }}
  labels:
    {{- include "[[ .Global.APPLICATION_NAME ]].labels" . | nindent 4 }}
spec:
  {{- if not .Values.autoscaling.enabled }}
  replicas: {{ .Values.replicaCount }}
  {{- end }}
  selector:
    matchLabels:
      {{- include "[[ .Global.APPLICATION_NAME ]].selectorLabels" . | nindent 6 }}
  template:
    metadata:
      {{- with .Values.podAnnotations }}
      annotations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      labels:
        {{- include "[[ .Global.APPLICATION_NAME ]].selectorLabels" . | nindent 8 }}
    spec:
      {{- with .Values.imagePullSecrets }}
      imagePullSecrets:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      serviceAccountName: {{ include "[[ .Global.APPLICATION_NAME ]].serviceAccountName" . }}
      securityContext:
        {{- toYaml .Values.podSecurityContext | nindent 8 }}
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.securityContext | nindent 12 }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          env:
            - name: SPRING_ACTIVE_PROFILE
              valueFrom:
                  configMapKeyRef:
                      key: SPRING_ACTIVE_PROFILE
                      name: {{ .Release.Name }}
          ports:
            - name: http
              containerPort: {{ .Values.service.targetport }}
              protocol: TCP
          readinessProbe:
            httpGet:
              path: {{ .Values.readiness.path }}
              port: {{ .Values.service.targetport }}
            initialDelaySeconds: {{ .Values.readiness.initialDelaySeconds }}
            periodSeconds: {{ .Values.readiness.periodSeconds }}
            failureThreshold: {{ .Values.readiness.failureThreshold }}
          #livenessProbe:
            #httpGet:
              #path: {{ .Values.liveness.path }}
              #port: {{ .Values.service.targetport }}
            #initialDelaySeconds: {{ .Values.liveness.initialDelaySeconds }}
            #periodSeconds: {{ .Values.liveness.periodSeconds }}
            #failureThreshold: {{ .Values.liveness.failureThreshold }}
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
