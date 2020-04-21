package global

import (
	"github.com/tech-showcase/entertainment-service/config"
	"github.com/tech-showcase/entertainment-service/presenter"
)

var Configuration = presenter.Config{}

func init() {
	var err error
	Configuration, err = config.Parse()
	if err != nil {
		panic(err)
	}
}
