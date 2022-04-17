/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/cocatrip/anchor/pkg/files"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "anchor",
	Short: "",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) { },
	// RunE: func(cmd *cobra.Command, args []string) error { },
}

func Execute() {
	files.LoadFiles()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config.yaml")
	}
	viper.SetConfigType("yaml")
}
