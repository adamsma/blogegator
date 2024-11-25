package config

import (
	"encoding/json"
	"log"
	"os"
)

func Read() (Config, error) {

	configPath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	cfgContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Unable to read config file: ", err)
	}

	var cfg Config
	err = json.Unmarshal(cfgContent, &cfg)
	if err != nil {
		log.Fatal("Error parsing config file: ", err)
	}

	return cfg, nil

}
