package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type (
	MovieListPerPage struct {
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

	MovieBlueprint struct {
		serverAddress string
		apiKey        string
	}
	MovieInterface interface {
		Search(string, int) (MovieListPerPage, error)
	}
)

func NewMovieModel(serverAddress string, apiKey string) MovieInterface {
	instance := MovieBlueprint{}
	instance.serverAddress = serverAddress
	instance.apiKey = apiKey

	return &instance
}

func (instance *MovieBlueprint) Search(keyword string, pageNumber int) (movies MovieListPerPage, err error) {
	req, err := http.NewRequest("GET", instance.serverAddress, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("apikey", instance.apiKey)
	q.Add("s", keyword)
	q.Add("page", strconv.Itoa(pageNumber))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.Unmarshal(respBody, &movies)
	if err != nil {
		return
	}

	return
}
