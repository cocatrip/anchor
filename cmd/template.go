/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	App          string
	BusinessName string
	Tag          string
	Template     string
	Jenkins      apps.Jenkins
	Helm         apps.Helm
	Docker       apps.Docker
}

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
		var c Config

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		c.Jenkins.New(c.App, c.Tag, c.BusinessName)

		template := c.Jenkins.TemplateJenkins()

		fmt.Println(template)

		saveFile(c.Jenkins.FILE, template)

		return nil
	},
}

var docker = &cobra.Command{
	Use:   "docker",
	Short: "parse docker from dockerfile",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c Config

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		c.Docker.New(c.App, c.Tag, c.BusinessName)

		template := c.Docker.TemplateDocker()

		fmt.Println(template)

		saveFile(c.Docker.FILE, template)

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
}

func saveFile(fileName string, content string) error {
	// check file ada ga
	if _, err := os.Stat(fileName); err != nil {
		// kalo gada create
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// save file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			panic(err)
		}
	} else {
		// kalo ada open
		file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// save file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
