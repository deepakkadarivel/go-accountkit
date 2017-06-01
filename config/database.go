package config

import (
	"fmt"
	"strings"
)

type Database struct {
	name             string
	connectionString string
}

func GetDatabaseConfig(configJSON *ConfigJSON) (*Database, error) {
	err := validate(configJSON)
	if err != nil {
		return nil, fmt.Errorf("DB Config initialization failed: %s.", err)
	}
	connectionString := buildConnectionString(configJSON)
	database := Database{configJSON.DBName, connectionString}
	return &database, nil
}

func (db *Database) ConnectionString() string {
	return db.connectionString
}

func validate(configJSON *ConfigJSON) error {
	missing := []string{}
	if configJSON.DBPort == 0 {
		missing = append(missing, "DB_PORT")
	}
	if configJSON.DBDriver == "" {
		missing = append(missing, "DB_DRIVER")
	}
	if configJSON.DBName == "" {
		missing = append(missing, "DB_NAME")
	}
	if configJSON.DBSchema == "" {
		missing = append(missing, "DB_SCHEMA")
	}
	if configJSON.DBUsername == "" {
		missing = append(missing, "DB_USERNAME")
	}

	if len(missing) > 0 {
		return fmt.Errorf("%s not found", strings.Join(missing, ", "))
	}

	return nil
}

func buildConnectionString(configJSON *ConfigJSON) string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s search_path=%s user=%s sslmode=disable",
		configJSON.DBServer,
		configJSON.DBPort,
		configJSON.DBName,
		configJSON.DBSchema,
		configJSON.DBUsername,
	)
}