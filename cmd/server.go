package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tech-showcase/entertainment-service/api"
	"github.com/tech-showcase/entertainment-service/helper"
)

type (
	serverFlagsStruct struct {
		Port int `json:"port"`
	}
)

var (
	serverFlags serverFlagsStruct

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run web server",
		Run: func(cmd *cobra.Command, args []string) {
			helper.RegistrarInstance.Register()
			defer helper.RegistrarInstance.Deregister()

			api.Activate(serverFlags.Port)
		},
	}
)

func init() {
	serverCmd.Flags().IntVarP(&serverFlags.Port, "port", "p", 8080, "Port which service will listen to")

	rootCmd.AddCommand(serverCmd)
}
