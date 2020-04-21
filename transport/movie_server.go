package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	movieProto "github.com/tech-showcase/entertainment-service/proto/movie"
)

type (
	movieGRPCServer struct {
		searchHandler grpctransport.Handler
	}
)

func NewMovieGRPCServer(ctx context.Context, endpoints MovieEndpoints) movieProto.MovieServer {
	instance := movieGRPCServer{}
	instance.searchHandler = grpctransport.NewServer(
		endpoints.SearchMovieEndpoint,
		DecodeSearchMovieRequest,
		EncodeSearchMovieResponse,
	)

	return &instance
}

func (instance *movieGRPCServer) Search(ctx context.Context, r *movieProto.SearchRequest) (*movieProto.SearchResponse, error) {
	_, resp, err := instance.searchHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*movieProto.SearchResponse), nil
}
