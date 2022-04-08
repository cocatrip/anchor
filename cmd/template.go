/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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

var templateCmd = &cobra.Command{
	Use:       "template",
	Short:     "Parse & save file",
	Long:      ``,
	ValidArgs: []string{"jenkins", "docker", "helm", "all"},
}

var jenkinsCmd = &cobra.Command{
	Use:          "jenkins",
	Short:        "Parse jenkins from jenkinsfile",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		f, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			return err
		}

		if err := yaml.Unmarshal(f, &c); err != nil {
			return err
		}

		apps.InitJenkins()

		resultFileName := fmt.Sprintf("Jenkinsfile-%s", c.Global["TESTING_TAG"])

		err = c.Template("Jenkinsfile", resultFileName)
		if err != nil {
			return err
		}

		return nil
	},
}

var dockerCmd = &cobra.Command{
	Use:          "docker",
	Short:        "Parse docker from dockerfile",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		f, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			return err
		}

		if err := yaml.Unmarshal(f, &c); err != nil {
			return err
		}

		apps.InitDocker()

		resultFileName := fmt.Sprintf("Dockerfile-%s", c.Global["TESTING_TAG"])

		err = c.Template("Dockerfile", resultFileName)
		if err != nil {
			return err
		}

		return nil
	},
}

var helmCmd = &cobra.Command{
	Use:          "helm",
	Short:        "Parse helm from values.yaml",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		f, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			return err
		}

		if err := yaml.Unmarshal(f, &c); err != nil {
			return err
		}

		isNoSecret, err := cmd.Flags().GetBool("no-secret")
		if err != nil {
			return err
		}

		templateFileName := "helm/values.yaml"
		resultFileName := fmt.Sprintf("helm/%s/values-%s.yaml", c.Global["APPLICATION_NAME"], c.Global["TESTING_TAG"])

		c.Helm["isNoSecret"] = isNoSecret

		_, err = os.Stat(templateFileName)
		if os.IsNotExist(err) {
			apps.InitHelm(c)
		}

		err = c.Template(templateFileName, resultFileName)
		if err != nil {
			return err
		}

		return nil
	},
}

var allCmd = &cobra.Command{
	Use:          "all",
	Short:        "Parse & save all (Jenkinsfile, Dockerfile, Helm)",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		f, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			return err
		}

		if err := yaml.Unmarshal(f, &c); err != nil {
			return err
		}

		// Jenkins
		apps.InitJenkins()
		jenkinsFileName := fmt.Sprintf("Jenkinsfile-%s", c.Global["TESTING_TAG"])
		err = c.Template("Jenkinsfile", jenkinsFileName)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println()

		// Docker
		apps.InitDocker()
		dockerFileName := fmt.Sprintf("Dockerfile-%s", c.Global["TESTING_TAG"])
		err = c.Template("Dockerfile", dockerFileName)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println()

		// Helm
		isNoSecret, err := cmd.Flags().GetBool("no-secret")
		if err != nil {
			return err
		}
		c.Helm["isNoSecret"] = isNoSecret
		err = apps.InitHelm(c)
		if err != nil {
			return err
		}
		helmTemplateFileName := "helm/values.yaml"
		helmResultFileName := fmt.Sprintf("helm/%s/values-%s.yaml", c.Global["APPLICATION_NAME"], c.Global["TESTING_TAG"])
		err = c.Template(helmTemplateFileName, helmResultFileName)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	// nambahin command template ke anchor
	rootCmd.AddCommand(templateCmd)
	// nambahin command jenkins ke template
	templateCmd.AddCommand(jenkinsCmd)
	// nambahin command docker ke template
	templateCmd.AddCommand(dockerCmd)
	// nambahin command helm ke template
	templateCmd.AddCommand(helmCmd)
	// nambahin command all ke template
	templateCmd.AddCommand(allCmd)

	helmCmd.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
	allCmd.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
}
