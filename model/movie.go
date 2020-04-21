package model

type (
	SearchMovieResponse struct {
		Response     string      `json:"Response"`
		Search       []MovieItem `json:"Search"`
		TotalResults string      `json:"totalResults"`
	}

	MovieItem struct {
		Poster string `json:"Poster"`
		Title  string `json:"Title"`
		Type   string `json:"Type"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
	}

	MovieBlueprint struct{}
	MovieInterface interface {
		Search(string, int) (SearchMovieResponse, error)
	}
)

func NewMovieModel() MovieInterface {
	instance := MovieBlueprint{}

	return &instance
}

func (instance *MovieBlueprint) Search(keyword string, pageNumber int) (movies SearchMovieResponse, err error) {
	return
}
