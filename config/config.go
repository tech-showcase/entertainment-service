package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		ServiceName string `json:"service_name"`
		Movie       Movie  `json:"movie"`
		Tracer      Tracer `json:"tracer"`
	}

	Movie struct {
		ServerAddress string `json:"server_address"`
		ApiKey        string `json:"api_key"`
	}

	Tracer struct {
		AgentAddress string `json:"agent_address"`
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

	viper.SetDefault("TRACER_AGENT_ADDRESS", "localhost")
	viper.BindEnv("TRACER_AGENT_ADDRESS")
}

func Parse() (config Config, err error) {
	configFilepath, configFilename := getConfigFile()
	viper.SetConfigName(configFilename)
	viper.SetConfigType("env")
	viper.AddConfigPath(configFilepath)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	config = getConfig()
	return
}

func getConfigFile() (configFilepath string, configFilename string) {
	configFilepath = viper.Get("CONFIG_FILEPATH").(string)
	configFilename = viper.Get("CONFIG_FILENAME").(string)
	return
}

func getConfig() (config Config) {
	config = Config{
		ServiceName: viper.Get("SERVICE_NAME").(string),
		Movie: Movie{
			ServerAddress: viper.Get("MOVIE_SERVER_ADDRESS").(string),
			ApiKey:        viper.Get("MOVIE_API_KEY").(string),
		},
		Tracer: Tracer{
			AgentAddress: viper.Get("TRACER_AGENT_ADDRESS").(string),
		},
	}
	return
}
