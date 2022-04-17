package cmd

import (
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:       "clean",
	Short:     "clean up files created by anchor",
	Long:      ``,
	ValidArgs: []string{"jenkins", "docker", "helm", "all"},
}

var cleanJenkinsCmd = &cobra.Command{
	Use:          "jenkins",
	Short:        "clean up jenkins related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var cleanDockerCmd = &cobra.Command{
	Use:          "docker",
	Short:        "clean up docker related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var cleanHelmCmd = &cobra.Command{
	Use:          "helm",
	Short:        "clean up helm related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var cleanAllCmd = &cobra.Command{
	Use:          "all",
	Short:        "clean up anchor related file",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.AddCommand(cleanJenkinsCmd)
	cleanCmd.AddCommand(cleanDockerCmd)
	cleanCmd.AddCommand(cleanHelmCmd)
	cleanCmd.AddCommand(cleanAllCmd)
}
