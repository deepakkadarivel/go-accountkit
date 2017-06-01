package config_test

import (
	"testing"
	"Accountkit/config"
	"github.com/stretchr/testify/assert"
	"fmt"
	"errors"
)

func TestGetDatabaseConfigForCorrectConfig(t *testing.T) {
	expectedConnectionString := "host=localhost port=5432 dbname=development search_path=accountkit user=deepakkv sslmode=disable"
	configJson := config.ConfigJSON{
		"postgres",
		5432,
		"localhost",
		"development",
		"accountkit",
		"deepakkv",
		"true",
	}
	actualDBConfig, err := config.GetDatabaseConfig(&configJson)
	assert.Nil(t, err)
	assert.Equal(t, expectedConnectionString, actualDBConfig.ConnectionString())
}

func TestGetDatabaseConfigForInvalidValuesProvidedForDBAttributes(t *testing.T) {
	configJSON := config.ConfigJSON{}
	expectedError := fmt.Errorf("DB Config initialization failed: %s not found.", errors.New("DB_PORT, DB_DRIVER, DB_NAME, DB_SCHEMA, DB_USERNAME"))
	actualDBConfig, err := config.GetDatabaseConfig(&configJSON)
	assert.Nil(t, actualDBConfig)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
