package movie

import (
	generalEndpoint "github.com/tech-showcase/entertainment-service/endpoint"
	"github.com/tech-showcase/entertainment-service/service"
)

type (
	Endpoint struct {
		Search generalEndpoint.GRPCEndpoint
	}
)

func NewMovieEndpoint(svc service.MovieService) Endpoint {
	instance := Endpoint{}
	instance.Search = generalEndpoint.GRPCEndpoint{
		Endpoint: makeSearchMovieEndpoint(svc),
		Decoder:  decodeSearchMovieRequest,
		Encoder:  encodeSearchMovieResponse,
	}

	return instance
}
