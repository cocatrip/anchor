package cmd

import (
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
		err := templateJenkins()
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
		err := templateDocker()
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
		// run function for templating helm
		err := templateHelm(cmd)
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
		// run function for templating jenkins
		err := templateJenkins()
		if err != nil {
			return err
		}

		// run function for templating docker
		err = templateDocker()
		if err != nil {
			return err
		}

		// run function for templating helm
		err = templateHelm(cmd)
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
