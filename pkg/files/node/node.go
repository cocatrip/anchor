package node

import (
	_ "embed"
)

//go:embed "configmap.yaml"
var ConfigMap string

//go:embed "deployment.yaml"
var Deployment string

//go:embed "Dockerfile"
var Dockerfile string

//go:embed "Jenkinsfile"
var Jenkinsfile string

//go:embed "secret.yaml"
var Secret string

//go:embed "service.yaml"
var Service string

//go:embed "values.yaml"
var Values string
