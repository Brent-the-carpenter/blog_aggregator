package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func getConfigFilePath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error finding user's home directory: %w ", err)
	}
	configFilePath := filepath.Join(userHome, configFileName)
	return configFilePath, nil
}

func (config *Config) SetUser(userName string) error {
	config.CurrentUserName = userName
	return write(*config)

}

func write(cfg Config) error {

	configFilePath, err := getConfigFilePath()

	if err != nil {
		return fmt.Errorf("Error finding user's home directory: %w", err)
	}

	file, err := os.Create(configFilePath)

	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
