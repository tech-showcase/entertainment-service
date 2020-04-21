package service

import (
	"github.com/tech-showcase/entertainment-service/model"
)

type (
	movieService struct{}
	MovieService interface {
		Search(model.MovieInterface, string, int) (model.MovieListPerPage, error)
	}
)

func NewMovieService() MovieService {
	instance := movieService{}

	return &instance
}

func (movieService) Search(movieModel model.MovieInterface, keyword string, pageNumber int) (movieData model.MovieListPerPage, err error) {
	movieData, err = movieModel.Search(keyword, pageNumber)
	if err != nil {
		return
	}

	return
}
