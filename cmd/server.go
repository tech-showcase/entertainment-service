package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	movieEndpoint "github.com/tech-showcase/entertainment-service/endpoint/movie"
	"github.com/tech-showcase/entertainment-service/global"
	"github.com/tech-showcase/entertainment-service/model"
	movieProto "github.com/tech-showcase/entertainment-service/proto/movie"
	"github.com/tech-showcase/entertainment-service/service"
	"github.com/tech-showcase/entertainment-service/transport"
	"google.golang.org/grpc"
	"net"
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
			config := global.Configuration.Movie
			movieModel := model.NewMovieModel(config.ServerAddress, config.ApiKey)
			movieService := service.NewMovieService(movieModel)
			movieEndpoints := movieEndpoint.NewMovieEndpoint(movieService)

			gRPCServer := grpc.NewServer()
			movieHandler := transport.NewMovieGRPCServer(movieEndpoints)
			movieProto.RegisterMovieServer(gRPCServer, movieHandler)

			portStr := fmt.Sprintf(":%d", serverFlags.Port)
			listener, err := net.Listen("tcp", portStr)
			if err != nil {
				panic(err)
			}

			err = gRPCServer.Serve(listener)
			if err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	serverCmd.Flags().IntVarP(&serverFlags.Port, "port", "p", 8080, "Port which service will listen to")

	rootCmd.AddCommand(serverCmd)
}
