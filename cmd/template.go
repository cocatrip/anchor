/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Parse & save file",
	Long:  ``,
}

var jenkins = &cobra.Command{
	Use:   "jenkins",
	Short: "Parse jenkins from jenkinsfile",
	Long:  ``,
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

		c.Template("Jenkinsfile", resultFileName)

		return nil
	},
}

var docker = &cobra.Command{
	Use:   "docker",
	Short: "Parse docker from dockerfile",
	Long:  ``,
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

		c.Template("Dockerfile", resultFileName)

		return nil
	},
}

var helm = &cobra.Command{
	Use:   "helm",
	Short: "Parse helm from values.yaml",
	Long:  ``,
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

		c.Helm["isNoSecret"] = isNoSecret
		apps.InitHelm(c)

		templateFileName := fmt.Sprintf("helm/%s/values.yaml", c.Global["APPLICATION_NAME"])
		resultFileName := fmt.Sprintf("helm/%s/values-%s.yaml", c.Global["APPLICATION_NAME"], c.Global["TESTING_TAG"])

		c.Template(templateFileName, resultFileName)

		return nil
	},
}

var all = &cobra.Command{
	Use:   "all",
	Short: "Parse & save all (Jenkinsfile, Dockerfile, Helm)",
	Long:  ``,
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
		c.Template("Jenkinsfile", jenkinsFileName)

	fmt.Println()
	fmt.Println()

		// Docker
		apps.InitDocker()
		dockerFileName := fmt.Sprintf("Dockerfile-%s", c.Global["TESTING_TAG"])
		c.Template("Dockerfile", dockerFileName)

	fmt.Println()
	fmt.Println()

		// Helm
		isNoSecret, err := cmd.Flags().GetBool("no-secret")
		if err != nil {
			return err
		}
		c.Helm["isNoSecret"] = isNoSecret
		apps.InitHelm(c)
		helmTemplateFileName := fmt.Sprintf("helm/%s/values.yaml", c.Global["APPLICATION_NAME"])
		helmResultFileName := fmt.Sprintf("helm/%s/values-%s.yaml", c.Global["APPLICATION_NAME"], c.Global["TESTING_TAG"])
		c.Template(helmTemplateFileName, helmResultFileName)

		return nil
	},
}

func init() {
	// nambahin command template ke anchor
	rootCmd.AddCommand(templateCmd)
	// nambahin command jenkins ke template
	templateCmd.AddCommand(jenkins)
	// nambahin command docker ke template
	templateCmd.AddCommand(docker)
	// nambahin command helm ke template
	templateCmd.AddCommand(helm)
	// nambahin command all ke template
	templateCmd.AddCommand(all)

	helm.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
	all.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
}
