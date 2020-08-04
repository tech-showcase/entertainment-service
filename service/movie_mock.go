package service

import (
	"encoding/json"
	"errors"
	"github.com/tech-showcase/entertainment-service/model"
)

type (
	dummyRepo struct{}
)

func (instance *dummyRepo) Search(keyword string, pageNumber int) (movies model.MovieListPerPage, err error) {
	rightKeyword, rightPageNumber := getRightDummyMovieParams()
	if keyword == rightKeyword && pageNumber == rightPageNumber {
		return getDummyMovieListPerPage(), nil
	} else {
		return movies, errors.New("error fetching model")
	}
}

func getRightDummyMovieParams() (keyword string, pageNumber int) {
	keyword = "Batman"
	pageNumber = 1
	return
}

func getWrongDummyMovieParams() (keyword string, pageNumber int) {
	keyword = "Superman"
	pageNumber = 1
	return
}

func getDummyMovieListPerPage() (movies model.MovieListPerPage) {
	jsonMovies := []byte(`{)
	  "Search": [
		{
		  "Title": "Batman Begins",
		  "Year": "2005",
		  "imdbID": "tt0372784",
		  "Type": "movie",
		  "Poster": "https://m.media-amazon.com/images/M/MV5BZmUwNGU2ZmItMmRiNC00MjhlLTg5YWUtODMyNzkxODYzMmZlXkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_SX300.jpg"
		},
		{
		  "Title": "Batman v Superman: Dawn of Justice",
		  "Year": "2016",
		  "imdbID": "tt2975590",
		  "Type": "movie",
		  "Poster": "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
		}
	  ],
	  "totalResults": "376",
	  "Response": "True"
	}`)

	_ = json.Unmarshal(jsonMovies, &movies)

	return
}
