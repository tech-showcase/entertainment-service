package service

import (
	"github.com/tech-showcase/entertainment-service/model"
)

type (
	movieService struct {
		movieRepo model.MovieRepo
	}
	MovieService interface {
		Search(string, int) (model.MovieListPerPage, error)
	}
)

func NewMovieService(movieRepo model.MovieRepo) MovieService {
	instance := movieService{}
	instance.movieRepo = movieRepo

	return &instance
}

func (instance *movieService) Search(keyword string, pageNumber int) (movieData model.MovieListPerPage, err error) {
	movieData, err = instance.movieRepo.Search(keyword, pageNumber)
	if err != nil {
		return
	}

	return
}
