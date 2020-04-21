package presenter

import "github.com/tech-showcase/entertainment-service/model"

type (
	SearchMovieRequest struct {
		Keyword    string `json:"keyword"`
		PageNumber int    `json:"page_number"`
	}
	SearchMovieResponse struct {
		model.MovieListPerPage
	}
)
