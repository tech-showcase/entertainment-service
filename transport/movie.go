package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	movieEndpoint "github.com/tech-showcase/entertainment-service/endpoint/movie"
	movieProto "github.com/tech-showcase/entertainment-service/proto/movie"
)

type (
	movieGRPCServer struct {
		searchHandler grpctransport.Handler
	}
)

func NewMovieGRPCServer(endpoint movieEndpoint.Endpoint) movieProto.MovieServer {
	instance := movieGRPCServer{}
	instance.searchHandler = grpctransport.NewServer(
		endpoint.Search.Endpoint,
		endpoint.Search.Decoder,
		endpoint.Search.Encoder)

	return &instance
}

func (instance *movieGRPCServer) Search(ctx context.Context, r *movieProto.SearchRequest) (*movieProto.SearchResponse, error) {
	_, resp, err := instance.searchHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*movieProto.SearchResponse), nil
}
