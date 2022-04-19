package files

import (
	_ "embed"

	"github.com/cocatrip/anchor/pkg/files/flutter"
	"github.com/cocatrip/anchor/pkg/files/maven"
	"github.com/cocatrip/anchor/pkg/files/node"
)

var Project string

var ConfigMap string
var Deployment string
var Dockerfile string
var Jenkinsfile string
var Secret string
var Service string
var Values string

// Load files according to project template
func LoadFiles() {
	if Project == "maven" {
		ConfigMap = maven.ConfigMap
		Deployment = maven.Deployment
		Dockerfile = maven.Dockerfile
		Jenkinsfile = maven.Jenkinsfile
		Secret = maven.Secret
		Service = maven.Service
		Values = maven.Values
	} else if Project == "node" {
		ConfigMap = node.ConfigMap
		Deployment = node.Deployment
		Dockerfile = node.Dockerfile
		Jenkinsfile = node.Jenkinsfile
		Secret = node.Secret
		Service = node.Service
		Values = node.Values
	} else if Project == "flutter" {
		ConfigMap = flutter.ConfigMap
		Deployment = flutter.Deployment
		Dockerfile = flutter.Dockerfile
		Jenkinsfile = flutter.Jenkinsfile
		Secret = flutter.Secret
		Service = flutter.Service
		Values = flutter.Values
	}
}
