package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		ServiceName string `json:"service_name"`
		Movie       Movie  `json:"movie"`
		Tracer      Tracer `json:"tracer"`
		Log         Log    `json:"log"`
		Consul      Consul `json:"consul"`
	}

	Movie struct {
		ServerAddress string `json:"server_address"`
		ApiKey        string `json:"api_key"`
	}

	Tracer struct {
		AgentAddress string `json:"agent_address"`
	}

	Log struct {
		Filepath string `json:"filepath"`
	}

	Consul struct {
		AgentAddress string        `json:"agent_address"`
		Service      ConsulService `json:"service"`
	}

	ConsulService struct {
		ID      string `json:"id"`
		Address string `json:"address"`
		Port    int    `json:"port"`
	}
)

var Instance Config

func init() {
	viper.SetDefault("CONFIG_FILEPATH", ".")
	viper.BindEnv("CONFIG_FILEPATH")
	viper.SetDefault("CONFIG_FILENAME", ".env")
	viper.BindEnv("CONFIG_FILENAME")

	viper.SetDefault("SERVICE_NAME", "entertainment-service")
	viper.BindEnv("SERVICE_NAME")

	viper.SetDefault("MOVIE_SERVER_ADDRESS", "http://localhost")
	viper.BindEnv("MOVIE_SERVER_ADDRESS")
	viper.SetDefault("MOVIE_API_KEY", "api-key")
	viper.BindEnv("MOVIE_API_KEY")

	viper.SetDefault("TRACER_AGENT_ADDRESS", "localhost:5775")
	viper.BindEnv("TRACER_AGENT_ADDRESS")

	viper.SetDefault("LOG_FILEPATH", "./server.log")
	viper.BindEnv("LOG_FILEPATH")

	viper.SetDefault("CONSUL_AGENT_ADDRESS", "localhost:8500")
	viper.BindEnv("CONSUL_AGENT_ADDRESS")
	viper.SetDefault("CONSUL_SERVICE_ID", "entertainment-service")
	viper.BindEnv("CONSUL_SERVICE_ID")
	viper.SetDefault("CONSUL_SERVICE_ADDRESS", "localhost")
	viper.BindEnv("CONSUL_SERVICE_ADDRESS")
	viper.SetDefault("CONSUL_SERVICE_PORT", "8080")
	viper.BindEnv("CONSUL_SERVICE_PORT")
}

func Parse() (config Config, err error) {
	configFilepath, configFilename := getConfigFile()
	viper.SetConfigName(configFilename)
	viper.SetConfigType("env")
	viper.AddConfigPath(configFilepath)

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	config = getConfig()
	return config, nil
}

func getConfigFile() (configFilepath string, configFilename string) {
	configFilepath = viper.GetString("CONFIG_FILEPATH")
	configFilename = viper.GetString("CONFIG_FILENAME")
	return
}

func getConfig() (config Config) {
	config = Config{
		ServiceName: viper.GetString("SERVICE_NAME"),
		Movie: Movie{
			ServerAddress: viper.GetString("MOVIE_SERVER_ADDRESS"),
			ApiKey:        viper.GetString("MOVIE_API_KEY"),
		},
		Tracer: Tracer{
			AgentAddress: viper.GetString("TRACER_AGENT_ADDRESS"),
		},
		Log: Log{
			Filepath: viper.GetString("LOG_FILEPATH"),
		},
		Consul: Consul{
			AgentAddress: viper.GetString("CONSUL_AGENT_ADDRESS"),
			Service: ConsulService{
				ID:      viper.GetString("CONSUL_SERVICE_ID"),
				Address: viper.GetString("CONSUL_SERVICE_ADDRESS"),
				Port:    viper.GetInt("CONSUL_SERVICE_PORT"),
			},
		},
	}
	return
}
