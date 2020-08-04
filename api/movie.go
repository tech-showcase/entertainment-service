package api

import (
	"github.com/tech-showcase/entertainment-service/config"
	movieEndpoint "github.com/tech-showcase/entertainment-service/endpoint/movie"
	"github.com/tech-showcase/entertainment-service/model"
	movieProto "github.com/tech-showcase/entertainment-service/proto/movie"
	"github.com/tech-showcase/entertainment-service/service"
	"github.com/tech-showcase/entertainment-service/transport"
	"google.golang.org/grpc"
)

func RegisterMovieGRPCAPI(gRPCServer *grpc.Server) {
	configuration := config.Instance
	movieModel := model.NewMovieModel(configuration.Movie.ServerAddress, configuration.Movie.ApiKey)
	movieService := service.NewMovieService(movieModel)
	movieEndpoints := movieEndpoint.NewMovieEndpoint(movieService)
	movieHandler := transport.NewMovieGRPCServer(movieEndpoints)
	movieProto.RegisterMovieServer(gRPCServer, movieHandler)
}
