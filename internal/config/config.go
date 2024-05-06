package config

import (
	"encoding/json"
	"log-parser/internal/projectpath"
	"log/slog"
	"os"
)

var Values Config

func Init(filePath string) {
	var err error

	Values, err = loadConfig(filePath)
	if err != nil {
		slog.Error("error", "err", err)
		os.Exit(1)
	}
}

func loadConfig(filePath string) (Config, error) {
	rootPath := projectpath.Root

	//open up absolute path to the config file
	file, err := os.Open(rootPath + "/" + filePath)

	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
