package main

import (
	"fmt"
	"github.com/tech-showcase/entertainment-service/cmd"
	"github.com/tech-showcase/entertainment-service/config"
)

func init() {
	var err error
	config.Instance, err = config.Parse()
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hi, I am " + config.Instance.ServiceName + "!")

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
