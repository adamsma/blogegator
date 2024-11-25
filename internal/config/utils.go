package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// use filepath.Join for system specific path separator
	return filepath.Join(home, configFileName), nil

}

func write(cfg Config) error {

	file, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
