/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "parse & save file",
	Long:  ``,
}

func Template(f string, m map[string]interface{}) {
    t := template.New(f)
    t = t.Delims("[[", "]]")
    t, err := t.ParseFiles(f)
    if err != nil {
        panic(err)
    }

    file, err := os.Create(f+"-uat")
    if err != nil {
        panic(err)
    }; defer file.Close()

    err = t.Execute(file, m)
    if err != nil {
        panic(err)
    }

    contentByte, err := ioutil.ReadFile(f+"-uat")
    if err != nil {
        panic(err)
    }; content := string(contentByte)

    fmt.Println(content)

    if strings.Contains(content, "<no value>") {
        fmt.Println("ada yang no value")
    } else {
        fmt.Println("tidak ada yang no value")
    }
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

        // Template("Jenkinsfile", c.Jenkins)

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

        fmt.Println(c.Docker)
        fmt.Println()

        Template("Dockerfile", c.Docker)

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
