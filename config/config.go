package config

import (
	"os"
	"fmt"
	"io"
	"encoding/json"
)

type ConfigJSON struct {
	DBDriver   string `json:"DB_DRIVER"`
	DBPort     int    `json:"DB_PORT"`
	DBServer   string `json:"DB_SERVER"`
	DBName     string `json:"DB_NAME"`
	DBSchema   string `json:"DB_SCHEMA"`
	DBUsername string `json:"DB_USERNAME"`
	DEVMode    string `json:"DEV_MODE"`
}

type Config struct {
	DB *Database
}

func NewAppSettingsFromFile(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %s", err)
	}

	//return DecodeJSON(file)
	var configJSON ConfigJSON
	err = json.NewDecoder(file).Decode(&configJSON)
	if err != nil {
		return nil, fmt.Errorf("Error processing config file: %s", err)
	}

	return generateConfigFromConfigJSON(&configJSON)
}

func DecodeJSON(reader io.Reader) (*Config, error) {
	var configJSON ConfigJSON
	err := json.NewDecoder(reader).Decode(&configJSON)
	if err != nil {
		return nil, fmt.Errorf("Error processing config file: %s", err)
	}

	return generateConfigFromConfigJSON(&configJSON)
}

func generateConfigFromConfigJSON(configJSON *ConfigJSON)  (*Config, error) {
	db, err := GetDatabaseConfig(configJSON)
	if err != nil {
		return nil, err
	}
	config := Config{DB: db}
	return &config, nil
}
