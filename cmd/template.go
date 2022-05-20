package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/cobra"
)

// represents acnhor template command
var templateCmd = &cobra.Command{
	Use:       "template",
	Short:     "Parse & save file",
	Long:      ``,
	ValidArgs: []string{"jenkins", "docker", "helm", "all"},
}

// represents anchor template jenkins command
var templateJenkinsCmd = &cobra.Command{
	Use:          "jenkins",
	Short:        "Parse jenkins from jenkinsfile",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// run function for templating jenkins
		err := template(cmd, "jenkins")
		if err != nil {
			return err
		}

		return nil
	},
}

// represents anchor template docker command
var templateDockerCmd = &cobra.Command{
	Use:          "docker",
	Short:        "Parse docker from dockerfile",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// run function for templating docker
		err := template(cmd, "docker")
		if err != nil {
			return err
		}

		return nil
	},
}

// represents anchor template helm command
var templateHelmCmd = &cobra.Command{
	Use:          "helm",
	Short:        "Parse helm from values.yaml",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := isHelmExist(); err != nil {
			return err
		}

		// run function for templating helm
		err := template(cmd, "helm")
		if err != nil {
			return err
		}

		return nil
	},
}

// represents anchor template all command
var templateAllCmd = &cobra.Command{
	Use:          "all",
	Short:        "Parse & save all (Jenkinsfile, Dockerfile, Helm)",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// run function for templating jenkins
		err = template(cmd, "jenkins")
		if err != nil {
			return err
		}

		// run function for templating docker
		err = template(cmd, "docker")
		if err != nil {
			return err
		}

		if err := isHelmExist(); err != nil {
			return err
		}

		// run function for templating helm
		err = template(cmd, "helm")
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	// add template command to anchor command
	rootCmd.AddCommand(templateCmd)

	// add jenkins command to template command
	templateCmd.AddCommand(templateJenkinsCmd)
	// add docker command to template command
	templateCmd.AddCommand(templateDockerCmd)
	// add helm command to template command
	templateCmd.AddCommand(templateHelmCmd)
	// add all command to template command
	templateCmd.AddCommand(templateAllCmd)

	// add --no-secret flag for anchor template helm command
	templateHelmCmd.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
	// add --no-secret flag for anchor template all command
	// which also contain helm command
	templateAllCmd.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
}

func isHelmExist() error {
	cmd := exec.Command("helm", "version")
  if err := cmd.Run(); err != nil {
		return err
  }
  return nil
}

func buildResultFileName(tools, toolFileName, testingTag, appName string) string {
	if tools != "helm" {
		return fmt.Sprintf("%s-%s", toolFileName, testingTag)
	} else {
		return fmt.Sprintf("helm/%s/values-%s.yaml", appName, testingTag)
	}
}

func template(cmd *cobra.Command, tools string) error {
	var err error
	var config apps.Config
	var toolFileName string
	var resultFileName string
	
	// read from config file and put it to config struct
	if err := ReadConfig(&config); err != nil {
		return err
	}

	testingTag := fmt.Sprintf("%s", config.Global["TESTING_TAG"])
	appName := fmt.Sprintf("%s", config.Global["APPLICATION_NAME"])

	if tools == "jenkins" {
		// call initJenkins to create template file for jenkins
		apps.InitJenkins()
		
		toolFileName = "Jenkinsfile"

	} else if tools == "docker" {
		// call initDocker to create template file for docker
		apps.InitDocker()

		toolFileName = "Dockerfile"

	} else if tools == "helm" {
		// read value of flag --no-secret
		isNoSecret, err := cmd.Flags().GetBool("no-secret")
		if err != nil {
			return err
		}

		// put it in config struct so it can be used for templating
		// example: [[ if .Helm.isNoSecret ]]
		config.Helm["isNoSecret"] = isNoSecret

		// define template file name
		toolFileName = "helm/values.yaml"

		// check if template file already exist
		_, err = os.Stat(toolFileName)
		if os.IsNotExist(err) {
			// if it doesn't then exec InitHelm to run `helm create APPLICATION_NAME`
			// where APPLICATION_NAME is defined in config.yaml
			apps.InitHelm(config)
		}
	}

	// define result file according to TESTING_TAG defined in config.yaml
	resultFileName = buildResultFileName(tools, toolFileName, testingTag, appName)

	// parse generated template and save output to resultFileName
	err = config.Template(toolFileName, resultFileName)
	if err != nil {
		return err
	}

	return nil
}
