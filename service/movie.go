package service

import (
	"github.com/tech-showcase/entertainment-service/model"
)

type (
	movieService struct {
		movieModel model.MovieInterface
	}
	MovieService interface {
		Search(string, int) (model.MovieListPerPage, error)
	}
)

func NewMovieService(movieModel model.MovieInterface) MovieService {
	instance := movieService{}
	instance.movieModel = movieModel

	return &instance
}

func (instance *movieService) Search(keyword string, pageNumber int) (movieData model.MovieListPerPage, err error) {
	movieData, err = instance.movieModel.Search(keyword, pageNumber)
	if err != nil {
		return
	}

	return
}
