package helper

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	stdconsul "github.com/hashicorp/consul/api"
	"github.com/tech-showcase/entertainment-service/config"
)

var RegistrarInstance *consul.Registrar

func NewConsulRegistrar(client consul.Client, serviceName string, serviceConfig config.ConsulService, logger log.Logger) *consul.Registrar {
	registration := &stdconsul.AgentServiceRegistration{
		ID:      serviceConfig.ID,
		Name:    serviceName,
		Tags:    []string{"go"},
		Port:    serviceConfig.Port,
		Address: serviceConfig.Address,
		Connect: &stdconsul.AgentServiceConnect{
			Native: false,
		},
	}

	registrar := consul.NewRegistrar(client, registration, log.With(logger, "component", "registrar"))
	return registrar
}

func NewConsulClient(agentAddress string) (consul.Client, error) {
	consulConfig := api.DefaultConfig()
	if len(agentAddress) > 0 {
		consulConfig.Address = agentAddress
	}

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}

	client := consul.NewClient(consulClient)
	return client, nil
}
