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

	movieRepo struct {
		serverAddress string
		apiKey        string
	}
	MovieRepo interface {
		Search(keyword string, pageNumber int) (movies MovieListPerPage, err error)
	}
)

func NewMovieModel(serverAddress string, apiKey string) MovieRepo {
	instance := movieRepo{}
	instance.serverAddress = serverAddress
	instance.apiKey = apiKey

	return &instance
}

func (instance *movieRepo) Search(keyword string, pageNumber int) (movies MovieListPerPage, err error) {
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
