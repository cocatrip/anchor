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
	Short: "parse & save file",
	Long:  ``,
}

var jenkins = &cobra.Command{
	Use:   "jenkins",
	Short: "parse jenkins from jenkinsfile",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		f, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(f, &c); err != nil {
			panic(err)
		}

		c.Template("Jenkinsfile", "Jenkinsfile-uat")

		return nil
	},
}

var docker = &cobra.Command{
	Use:   "docker",
	Short: "parse docker from dockerfile",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		f, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(f, &c); err != nil {
			panic(err)
		}

		c.Template("Dockerfile", "Dockerfile-uat")

		return nil
	},
}

var helm = &cobra.Command{
	Use:   "helm",
	Short: "parse helm from values.yaml",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		f, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(f, &c); err != nil {
			panic(err)
		}

		isNoSecret, err := cmd.Flags().GetBool("no-secret")
		if err != nil {
			panic(err)
		}

		apps.InitHelm(c, isNoSecret)

		templateFileName := fmt.Sprintf("helm/%s/values.yaml", c.Global["APPLICATION_NAME"])
		resultFileName := fmt.Sprintf("helm/%s/values-%s.yaml", c.Global["APPLICATION_NAME"], c.Global["TESTING_TAG"])

		c.Template(templateFileName, resultFileName)

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

	helm.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
}
