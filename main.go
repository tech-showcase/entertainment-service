package main

import (
	"fmt"
	"github.com/tech-showcase/entertainment-service/cmd"
	"github.com/tech-showcase/entertainment-service/config"
	"github.com/tech-showcase/entertainment-service/helper"
)

func init() {
	var err error
	config.Instance, err = config.Parse()
	if err != nil {
		panic(err)
	}

	//helper.LoggerInstance = helper.NewLogger()
	helper.LoggerInstance, err = helper.NewFileLogger(config.Instance.Log.Filepath)
	if err != nil {
		panic(err)
	}

	helper.TracerInstance, _, err = helper.NewTracer(config.Instance.ServiceName, config.Instance.Tracer.AgentAddress)
	if err != nil {
		helper.LoggerInstance.Log("NewTracer", err)
	}
}

func main() {
	fmt.Println("Hi, I am " + config.Instance.ServiceName + "!")

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
