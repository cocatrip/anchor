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
  Jenkins apps.Jenkins
	Helm 		apps.Helm
	Docker	apps.Docker
}

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:		"template",
	Short:	"",
	Long:		``,
}

var jenkins = &cobra.Command{
	Use:		"jenkins",
	Short:	"",
	Long:		``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c Config

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		template := c.Jenkins.TemplateJenkins()

		fmt.Println(template)

		saveFile("Jenkinsfile-"+"test", template)

		return nil
	},
}

var helm = &cobra.Command{
	Use:		"helm",
	Short:	"",
	Long:		``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c Config

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		template := c.Helm.TemplateHelm()
		fmt.Println(template)

		return nil
	},
}

var docker = &cobra.Command{
	Use:		"docker",
	Short:	"",
	Long:		``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c Config

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		template := c.Docker.TemplateDocker()
		fmt.Println(template)

		return nil
	},
}

func init() {
	// nambahin command template ke anchor
	rootCmd.AddCommand(templateCmd)
	// nambahin command jenkins ke template
	templateCmd.AddCommand(jenkins)
	// nambahin command helm ke template
	templateCmd.AddCommand(helm)
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
    }; defer file.Close()

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
		}; defer file.Close()

		// save file
		_, err = file.WriteAt([]byte(content), 0)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
