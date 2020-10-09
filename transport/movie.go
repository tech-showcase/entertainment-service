package transport

import (
	"context"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	movieEndpoint "github.com/tech-showcase/entertainment-service/endpoint/movie"
	"github.com/tech-showcase/entertainment-service/helper"
	movieProto "github.com/tech-showcase/entertainment-service/proto/movie"
)

type (
	movieGRPCServer struct {
		searchHandler grpctransport.Handler
	}
)

func NewMovieGRPCServer(endpoint movieEndpoint.Endpoint, logger log.Logger, tracer stdopentracing.Tracer) movieProto.MovieServer {
	var options []grpctransport.ServerOption
	options = append(options, grpctransport.ServerBefore(helper.GRPCToContext(tracer, "searchMovie-transport", logger)))

	instance := movieGRPCServer{}
	instance.searchHandler = grpctransport.NewServer(
		endpoint.Search.Endpoint,
		endpoint.Search.Decoder,
		endpoint.Search.Encoder,
		options...,
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
