package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {

	home, err := os.UserHomeDir()
	return home + "/" + configFileName, err

}

func write(cfg Config) error {

	file, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	path, _ := getConfigFilePath()
	wrtErr := os.WriteFile(path, file, 0644)
	if err != nil {
		return wrtErr
	}

	return nil
}
