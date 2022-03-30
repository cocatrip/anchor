/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/cocatrip/anchor/cmd/apps"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
  Jenkins apps.Jenkins
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
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
			fmt.Println()
		}

		var c Config

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		template := c.Jenkins.TemplateJenkins()
		fmt.Println(template)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.AddCommand(jenkins)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
