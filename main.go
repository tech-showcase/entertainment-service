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

	consulClient, err := helper.NewConsulClient(config.Instance.Consul.AgentAddress)
	if err != nil {
		panic(err)
	}
	helper.RegistrarInstance = helper.NewConsulRegistrar(
		consulClient,
		config.Instance.ServiceName,
		config.Instance.Consul.Service,
		helper.LoggerInstance,
	)
}

func main() {
	fmt.Println("Hi, I am " + config.Instance.ServiceName + "!")

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
