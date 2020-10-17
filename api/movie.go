package api

import (
	"github.com/tech-showcase/entertainment-service/config"
	movieEndpoint "github.com/tech-showcase/entertainment-service/endpoint/movie"
	"github.com/tech-showcase/entertainment-service/helper"
	"github.com/tech-showcase/entertainment-service/model"
	movieProto "github.com/tech-showcase/entertainment-service/proto/movie"
	"github.com/tech-showcase/entertainment-service/service"
	"github.com/tech-showcase/entertainment-service/transport"
	"google.golang.org/grpc"
)

func RegisterMovieGRPCAPI(gRPCServer *grpc.Server) {
	configuration := config.Instance
	loggerInstance := helper.LoggerInstance
	tracerInstance := helper.TracerInstance

	movieModel := model.NewMovieModel(configuration.Movie.ServerAddress, configuration.Movie.ApiKey)
	movieService := service.NewMovieService(movieModel)
	movieEndpoints := movieEndpoint.NewMovieEndpoint(movieService, loggerInstance, tracerInstance)
	movieHandler := transport.NewMovieGRPCServer(movieEndpoints, loggerInstance, tracerInstance)
	movieProto.RegisterMovieServer(gRPCServer, movieHandler)
}
