package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Reads config file in the home directory
func Read() (Config, error) {
	// Get config filepath
	configFilepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Open config file
	file, err := os.Open(configFilepath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// Decode file contents into Config struct
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// Set user for gator
func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

// Get filepath for config
func getConfigFilePath() (string, error) {
	homepath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homepath, configFileName), nil
}

// Write to config file
func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Create config file using the filepath
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode config struct to config file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
