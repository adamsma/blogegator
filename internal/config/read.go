package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {

	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	cfgContent, err := os.ReadFile(configPath)
	if err != nil {
		err = fmt.Errorf("unable to read config file: %v", err)
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(cfgContent, &cfg)
	if err != nil {
		err = fmt.Errorf("Error parsing config file: %v", err)
		return Config{}, err
	}

	return cfg, nil

}
