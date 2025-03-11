package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file: %v", err)
	}
	jsonBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %v", err)
	}
	var jsonStruct Config
	err = json.Unmarshal(jsonBytes, &jsonStruct)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling json file: %v", err)
	}
	return jsonStruct, nil
}

func (cfg Config) SetUser(user string) error {
	cfg.Current_user_name = user
	err := write(cfg)
	if err != nil {
		return fmt.Errorf("error writing to config file in home directory: %v", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error fetching home directory: %v", err)
	}
	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error fetching config file from home directory: %v", err)
	}
	file, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("error opening config file from home directory: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("error writing to config file: %v", err)
	}

	return nil
}
