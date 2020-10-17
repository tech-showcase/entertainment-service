package movie

import (
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	generalEndpoint "github.com/tech-showcase/entertainment-service/endpoint"
	"github.com/tech-showcase/entertainment-service/middleware"
	"github.com/tech-showcase/entertainment-service/service"
)

type (
	Endpoint struct {
		Search generalEndpoint.GRPCEndpoint
	}
)

func NewMovieEndpoint(svc service.MovieService, logger log.Logger, tracer stdopentracing.Tracer) Endpoint {
	movieEndpoint := Endpoint{}

	searchMovieEndpoint := makeSearchMovieEndpoint(svc)
	searchMovieEndpoint = middleware.ApplyTracerClient("searchMovie-endpoint", searchMovieEndpoint, tracer)
	searchMovieEndpoint = middleware.ApplyLogger("searchMovie", searchMovieEndpoint, logger)
	movieEndpoint.Search = generalEndpoint.GRPCEndpoint{
		Endpoint: searchMovieEndpoint,
		Decoder:  decodeSearchMovieRequest,
		Encoder:  encodeSearchMovieResponse,
	}

	return movieEndpoint
}
