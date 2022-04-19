package cmd

import (
	"github.com/spf13/cobra"
)

// cleanCmd represents the anchor clean command
var cleanCmd = &cobra.Command{
	Use:       "clean",
	Short:     "clean up files created by anchor",
	Long:      ``,
	ValidArgs: []string{"jenkins", "docker", "helm", "all"},
}

// cleanCmd represents the anchor clean jenkins command
var cleanJenkinsCmd = &cobra.Command{
	Use:          "jenkins",
	Short:        "clean up jenkins related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// run the corresponding function for this command
		err := cleanJenkins()
		if err != nil {
			return err
		}

		return nil
	},
}

// cleanCmd represents the anchor clean docker command
var cleanDockerCmd = &cobra.Command{
	Use:          "docker",
	Short:        "clean up docker related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// run the corresponding function for this command
		err := cleanDocker()
		if err != nil {
			return err
		}

		return nil
	},
}

// cleanCmd represents the anchor clean helm command
var cleanHelmCmd = &cobra.Command{
	Use:          "helm",
	Short:        "clean up helm related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// run the corresponding function for this command
		err := cleanHelm()
		if err != nil {
			return err
		}

		return err
	},
}

var cleanAllCmd = &cobra.Command{
	Use:          "all",
	Short:        "clean up anchor related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// run the corresponding function for this command
		err := cleanJenkins()
		if err != nil {
			return err
		}

		err = cleanDocker()
		if err != nil {
			return err
		}

		err = cleanHelm()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	// add clean command to anchor command
	rootCmd.AddCommand(cleanCmd)

	// add jenkins command to clean command
	cleanCmd.AddCommand(cleanJenkinsCmd)
	// add docker command to clean command
	cleanCmd.AddCommand(cleanDockerCmd)
	// add helm command to clean command
	cleanCmd.AddCommand(cleanHelmCmd)
	// add all command to clean command
	cleanCmd.AddCommand(cleanAllCmd)
}
