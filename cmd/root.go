package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "entertainment-service",
	Short: "Entertainment Service is a service that provide data about entertainment",
	Long: `A service that can be a portal for all entertainment data
				like movie, music, etc.`,
}

func Execute() error {
	return rootCmd.Execute()
}
