package presenter

type (
	Config struct {
		Movie `json:"movie"`
	}

	Movie struct {
		ServerAddress string `json:"server_address"`
		ApiKey        string `json:"api_key"`
	}
)
