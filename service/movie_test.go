package service

import (
	"encoding/json"
	"errors"
	"github.com/tech-showcase/entertainment-service/model"
	"reflect"
	"testing"
)

type (
	dummyModel struct{}
)

func (instance *dummyModel) Search(keyword string, pageNumber int) (movies model.MovieListPerPage, err error) {
	if keyword == "Batman" && pageNumber == 1 {
		return getDummyMovieListPerPage(), nil
	} else {
		return movies, errors.New("error fetching model")
	}
}

func getDummyMovieListPerPage() (movieData model.MovieListPerPage) {
	jsonData := []byte(`{)
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

	_ = json.Unmarshal(jsonData, &movieData)

	return
}

func TestMovieService_Search(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		expectedOutput := getDummyMovieListPerPage()

		dummyModel := dummyModel{}
		movieService := NewMovieService(&dummyModel)

		movieData, err := movieService.Search("Batman", 1)

		if err != nil {
			t.Fatal("an error has occurred")
		} else if !reflect.DeepEqual(movieData, expectedOutput) {
			t.Fatal("unexpected output")
		}
	})
	t.Run("negative", func(t *testing.T) {
		dummyModel := dummyModel{}
		movieService := NewMovieService(&dummyModel)

		_, err := movieService.Search("Batman", 2)

		if err == nil {
			t.Fatal("an error should occur")
		}
	})
}
