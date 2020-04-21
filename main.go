package main

import (
	"fmt"
	"github.com/tech-showcase/entertainment-service/cmd"
	"github.com/tech-showcase/entertainment-service/global"
	"github.com/tech-showcase/entertainment-service/model"
	movieProto "github.com/tech-showcase/entertainment-service/proto/movie"
	"github.com/tech-showcase/entertainment-service/service"
	"github.com/tech-showcase/entertainment-service/transport"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

func main() {
	fmt.Println("Hi, I am Entertainment Service!")

	args := cmd.Parse()

	ctx := context.Background()

	config := global.Configuration.Movie
	movieModel := model.NewMovieModel(config.ServerAddress, config.ApiKey)
	movieService := service.NewMovieService(movieModel)
	movieEndpoints := transport.MovieEndpoints{
		SearchMovieEndpoint: transport.MakeSearchMovieEndpoint(movieService),
	}

	gRPCServer := grpc.NewServer()
	movieHandler := transport.NewMovieGRPCServer(ctx, movieEndpoints)
	movieProto.RegisterMovieServer(gRPCServer, movieHandler)

	portStr := fmt.Sprintf(":%d", args.Port)
	listener, err := net.Listen("tcp", portStr)
	if err != nil {
		panic(err)
	}

	err = gRPCServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
