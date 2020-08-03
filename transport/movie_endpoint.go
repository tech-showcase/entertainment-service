package transport

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/tech-showcase/entertainment-service/presenter"
	movieProto "github.com/tech-showcase/entertainment-service/proto/movie"
	"github.com/tech-showcase/entertainment-service/service"
)

type MovieEndpoints struct {
	SearchMovieEndpoint endpoint.Endpoint
}

func MakeSearchMovieEndpoint(svc service.MovieService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(presenter.SearchMovieRequest)

		movieData, err := svc.Search(req.Keyword, req.PageNumber)
		if err != nil {
			return presenter.SearchMovieResponse{}, nil
		}
		return presenter.SearchMovieResponse{MovieListPerPage: movieData}, nil
	}
}

func DecodeSearchMovieRequest(_ context.Context, r interface{}) (interface{}, error) {
	if req, ok := r.(*movieProto.SearchRequest); ok {
		return presenter.SearchMovieRequest{
			Keyword:    req.Keyword,
			PageNumber: int(req.PageNumber),
		}, nil
	} else {
		return nil, errors.New("format request is wrong")
	}
}

func EncodeSearchMovieResponse(_ context.Context, r interface{}) (interface{}, error) {
	if res, ok := r.(presenter.SearchMovieResponse); ok {
		movies := make([]*movieProto.SearchResponse_MovieItem, 0)
		for _, item := range res.Search {
			movie := movieProto.SearchResponse_MovieItem{
				Poster: item.Poster,
				Title:  item.Title,
				Type:   item.Type,
				Year:   item.Year,
				ImdbId: item.ImdbID,
			}
			movies = append(movies, &movie)
		}

		return &movieProto.SearchResponse{
			Response:     res.Response,
			Search:       movies,
			TotalResults: res.TotalResults,
		}, nil
	} else {
		return nil, errors.New("format response is wrong")
	}
}
