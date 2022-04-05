package files

var valuesSecret string = `# Default values for ad1-lead-main.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: [[ .Global.SERVER_NAME ]]/[[ .Global.BUSINESS_NAME ]]/[[ .Global.TESTING_TAG ]]-[[ .Global.APPLICATION_NAME ]]
  pullPolicy: Always
  # Overrides the image tag whose default is the chart appVersion.
  tag: [[ Helm.Version_Major}.[[ Helm.Version_Minor ]].[[ Helm.Version_Patch ]]-[[ Helm.BUILD_TIMESTAMP ]]-[[ Helm.BUILD_NUMBER ]]

config:
  spring_active_profile: '[[ .Global.TESTING_TAG ]]'
  secret_name: '[[ .Helm.SECRETDB_NAME ]]'
  DB_USER: [[ .Helm.DB_USER ]]
  DB_PASSWORD: [[ .Helm.DB_PASSWORD ]]

readiness:
  path: /actuator/health/readiness
  failureThreshold: 3
  initialDelaySeconds: 10
  periodSeconds: 10

liveness:
  path: /actuator/health/liveness
  failureThreshold: 3
  initialDelaySeconds: 10
  periodSeconds: 10

imagePullSecrets:
  - name: regcred
nameOverride: ""
fullnameOverride: ""

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

podAnnotations: {}

podSecurityContext: {}
  # fsGroup: 2000

chart:
  releaseName: [[ .Global.APPLICATION_NAME ]]

securityContext: {}
  # capabilities:
  #   drop:
  #   - ALL
  # readOnlyRootFilesystem: true
  # runAsNonRoot: true
  # runAsUser: 1000

service:
  type: [[ .Helm.SERVICE_TYPE ]]
  port: [[ .Helm.SERVICE_PORT ]]
  targetport: [[ Helm.SERVICE_TARGETPORT ]]

ingress:
  enabled: false
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  hosts:
    - host: chart-example.local
      paths: []
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - chart-example.local

resources:
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  limits:
    cpu: [[ .Helm.CPU_LIMIT ]]
    memory: [[ .Helm.MEM_LIMIT ]]
  requests:
    cpu: [[ .Helm.CPU_REQUEST ]]
    memory: [[ .Helm.MEM_REQUEST ]]

autoscaling:
  enabled: true
  minReplicas: 1
  maxReplicas: 2
  targetCPUUtilizationPercentage: 80
  # targetMemoryUtilizationPercentage: 80

nodeSelector: {}

tolerations: []

affinity: {}`
