/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "",
	Long: ``,
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
		
		file, err := os.OpenFile(c.Jenkins.File, os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}; defer file.Close()

		_, err = file.WriteAt([]byte(template), 0)
		if err != nil {
			panic(err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
