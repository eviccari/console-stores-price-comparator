package configs

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	FileConfigName = "application"
	FileConfigType = "yml"
	BasePath       = "."
)

var Params struct {
	XboxStore struct {
		StoreBaseURI string `yaml:"storeBaseURI"`
	} `yaml:"xboxStore"`
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Load() {
	fileName := fmt.Sprintf("%s.%s", FileConfigName, FileConfigType)
	file, err := os.Open(fileName)
	if err != nil {
		logError(err)
		os.Exit(1)
	}
	err = yaml.NewDecoder(file).Decode(&Params)
	if err != nil {
		logError(err)
	}
}

func logError(err error) {
	logger.Error("load_configs", "error", err.Error())
}
