package cmd

import (
	"fmt"
	"os"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/cobra"
)

func clean(tools string) error {
	var config apps.Config
	var resultFileName string
	var removeList []string

	// read from config file and put it to config struct
	if err := ReadConfig(&config); err != nil {
		return err
	}

	if tools != "helm" {
		if tools == "jenkins" {
			// define the resultFileName so it can be deleted
			resultFileName = fmt.Sprintf("Jenkinsfile-%s", config.Global["TESTING_TAG"])

			// list of to be removed files
			removeList = []string{
				"Jenkinsfile",
				resultFileName,
			}
		} else if tools == "docker" {
			// define the resultFileName so it can be deleted
			resultFileName = fmt.Sprintf("Dockerfile-%s", config.Global["TESTING_TAG"])

			// list of to be removed files
			removeList = []string{
				"Dockerfile",
				resultFileName,
			}
		}

		// loop through the list and delete file one by one
		for _, v := range removeList {
			// check if exist if not then don't delete so os.Remove will never return err
			_, err := os.Stat(v)
			if !os.IsNotExist(err) {
				os.Remove(v)
				fmt.Printf("Removing %s\n", v)
			}
		}
	} else {
		// simply check if helm directory exist or not
		_, err := os.Stat("helm")
		// if exist then recursively remove the directory
		if !os.IsNotExist(err) {
			os.RemoveAll("helm")
			fmt.Printf("Removing helm directory\n")
		}
	}

	return nil
}

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
		err := clean("jenkins")
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
		err := clean("docker")
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
		err := clean("helm")
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
		err := clean("jenkins")
		if err != nil {
			return err
		}

		err = clean("docker")
		if err != nil {
			return err
		}

		err = clean("helm")
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
