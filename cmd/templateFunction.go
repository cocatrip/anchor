package cmd

import (
	"fmt"
	"os"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/cobra"
)

// function for 'anchor template jenkins'
func templateJenkins() error {
	var config apps.Config

	// read from config file and put it to config struct
	if err := ReadConfig(&config); err != nil {
		return err
	}

	// call initJenkins to create template file for jenkins
	apps.InitJenkins()

	// define result file according to TESTING_TAG defined in config.yaml
	resultFileName := fmt.Sprintf("Jenkinsfile-%s", config.Global["TESTING_TAG"])

	// parse generated template and save output to resultFileName
	err := config.Template("Jenkinsfile", resultFileName)
	if err != nil {
		return err
	}

	return nil
}

// function for 'anchor template docker'
func templateDocker() error {
	var config apps.Config

	// read from config file and put it to config struct
	if err := ReadConfig(&config); err != nil {
		return err
	}

	// call initDocker to create template file for docker
	apps.InitDocker()

	// define result file according to TESTING_TAG defined in config.yaml
	resultFileName := fmt.Sprintf("Dockerfile-%s", config.Global["TESTING_TAG"])

	// parse generated template and save output to resultFileName
	err := config.Template("Dockerfile", resultFileName)
	if err != nil {
		return err
	}

	return nil
}

// function for 'anchor template helm'
func templateHelm(cmd *cobra.Command) error {
	var config apps.Config

	// read from config file and put it to config struct
	if err := ReadConfig(&config); err != nil {
		return err
	}

	// read value of flag --no-secret
	isNoSecret, err := cmd.Flags().GetBool("no-secret")
	if err != nil {
		return err
	}
	// put it in config struct so it can be used for templating
	// example: [[ if .Helm.isNoSecret ]]
	config.Helm["isNoSecret"] = isNoSecret

	// define template file name
	templateFileName := "helm/values.yaml"
	// define result file according to TESTING_TAG defined in config.yaml
	resultFileName := fmt.Sprintf("helm/%s/values-%s.yaml", config.Global["APPLICATION_NAME"], config.Global["TESTING_TAG"])

	// check if template file already exist
	_, err = os.Stat(templateFileName)
	// if it doesn't then exec InitHelm to run `helm create APPLICATION_NAME`
	// where APPLICATION_NAME is defined in config.yaml
	if os.IsNotExist(err) {
		apps.InitHelm(config)
	}

	// parse generated template and save output to resultFileName
	err = config.Template(templateFileName, resultFileName)
	if err != nil {
		return err
	}

	return nil
}
