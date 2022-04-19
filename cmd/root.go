package cmd

import (
	"os"

	"github.com/cocatrip/anchor/pkg/files"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// represents anchor command
var rootCmd = &cobra.Command{
	Use:   "anchor",
	Short: "",
	Long:  ``,
}

// Execute function called by main.go
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init function executed before running any cli command
func init() {
	// run initConfig function
	cobra.OnInitialize(initConfig)

	// declare flag for anchor cli
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file")
	rootCmd.PersistentFlags().StringVarP(&files.Project, "type", "t", "maven", "maven, node, or flutter")
}

func initConfig() {
	// set config or default to config.yaml
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config.yaml")
	}
	viper.SetConfigType("yaml")

	// load all files according to project type
	files.LoadFiles()
}
