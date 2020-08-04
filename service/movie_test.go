package service

import (
	"reflect"
	"testing"
)

func TestMovieService_Search(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		expectedOutput := getDummyMovieListPerPage()

		dummyRepo := dummyRepo{}
		movieService := NewMovieService(&dummyRepo)

		movieData, err := movieService.Search(getRightDummyMovieParams())

		if err != nil {
			t.Fatal("an error has occurred")
		} else if !reflect.DeepEqual(movieData, expectedOutput) {
			t.Fatal("unexpected output")
		}
	})
	t.Run("negative", func(t *testing.T) {
		dummyModel := dummyRepo{}
		movieService := NewMovieService(&dummyModel)

		_, err := movieService.Search(getWrongDummyMovieParams())

		if err == nil {
			t.Fatal("an error should occur")
		}
	})
}
