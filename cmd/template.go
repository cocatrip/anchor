/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/cocatrip/anchor/pkg/files"
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

var templateJenkinsCmd = &cobra.Command{
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

var templateDockerCmd = &cobra.Command{
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

var templateHelmCmd = &cobra.Command{
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

var templateAllCmd = &cobra.Command{
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
	rootCmd.AddCommand(templateCmd)

	templateCmd.AddCommand(templateJenkinsCmd)
	templateCmd.AddCommand(templateDockerCmd)
	templateCmd.AddCommand(templateHelmCmd)
	templateCmd.AddCommand(templateAllCmd)

	templateCmd.Flags().StringVarP(&files.Project, "type", "t", "maven", "maven, node, or flutter")
	templateHelmCmd.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
	templateAllCmd.Flags().BoolP("no-secret", "", false, "don't create secret.yaml inside templates")
}
