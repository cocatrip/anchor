/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/cocatrip/anchor/pkg/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


// templateCmd represents the template command
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

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		c.Jenkins.New(c)

		template := c.Jenkins.TemplateJenkins()

		fmt.Println(template)

		common.SaveFile(c.Jenkins.FILE, template)

		return nil
	},
}

var docker = &cobra.Command{
	Use:   "docker",
	Short: "parse docker from dockerfile",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		c.Docker.New(c)

		template := c.Docker.TemplateDocker()

		fmt.Println(template)

		common.SaveFile(c.Docker.FILE, template)

		return nil
	},
}

var helm = &cobra.Command{
	Use:   "helm",
	Short: "parse helm from values.yaml",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c apps.Config

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		c.Helm.New(c)

		c.Helm.InitHelm()

		template := c.Helm.TemplateHelm()

		fmt.Println(template)

		// saveFile(c.Helm.FILE, template)

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
}
