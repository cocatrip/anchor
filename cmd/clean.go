package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:       "clean",
	Short:     "clean up files created by anchor",
	Long:      ``,
	ValidArgs: []string{"jenkins", "docker", "helm", "all"},
}

func cleanJenkins() error {
	var c apps.Config

	f, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(f, &c); err != nil {
		return err
	}

	resultFileName := fmt.Sprintf("Jenkinsfile-%s", c.Global["TESTING_TAG"])

	removeList := []string{
		"Jenkinsfile",
		resultFileName,
	}

	for _, v := range removeList {
		_, err := os.Stat(v)
		if !os.IsNotExist(err) {
			os.RemoveAll(v)
			fmt.Printf("Removing %s\n", v)
		}
	}

	return nil
}

var cleanJenkinsCmd = &cobra.Command{
	Use:          "jenkins",
	Short:        "clean up jenkins related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cleanJenkins()
		if err != nil {
			return err
		}

		return nil
	},
}

func cleanDocker() error {
	var c apps.Config

	f, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(f, &c); err != nil {
		return err
	}

	resultFileName := fmt.Sprintf("Dockerfile-%s", c.Global["TESTING_TAG"])

	removeList := []string{
		"Dockerfile",
		resultFileName,
	}

	for _, v := range removeList {
		_, err := os.Stat(v)
		if !os.IsNotExist(err) {
			os.RemoveAll(v)
			fmt.Printf("Removing %s\n", v)
		}
	}

	return nil
}

var cleanDockerCmd = &cobra.Command{
	Use:          "docker",
	Short:        "clean up docker related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cleanDocker()
		if err != nil {
			return err
		}

		return nil
	},
}

func cleanHelm() error {
	var c apps.Config

	f, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(f, &c); err != nil {
		return err
	}

	_, err = os.Stat("helm")
	if !os.IsNotExist(err) {
		os.RemoveAll("helm")
		fmt.Printf("Removing helm directory\n")
	}

	return nil
}

var cleanHelmCmd = &cobra.Command{
	Use:          "helm",
	Short:        "clean up helm related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
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
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.AddCommand(cleanJenkinsCmd)
	cleanCmd.AddCommand(cleanDockerCmd)
	cleanCmd.AddCommand(cleanHelmCmd)
	cleanCmd.AddCommand(cleanAllCmd)
}
