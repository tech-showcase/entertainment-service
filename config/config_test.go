package config

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedOutput := Config{
		Movie: Movie{
			ServerAddress: "http://dummy.address/",
			ApiKey:        "dummy-api-key",
		},
	}

	os.Setenv("CONFIG_FILEPATH", "..")
	os.Setenv("CONFIG_FILENAME", ".env.example")

	config, err := Parse()

	if err != nil {
		t.Fatal("an error has occurred")
	} else if !reflect.DeepEqual(config, expectedOutput) {
		t.Fatal("unexpected output")
	}
}
