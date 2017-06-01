package config_test

import (
	"testing"
	"Accountkit/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAppSettingsFromFile(t *testing.T) {
	configFilePath := "app.conf"
	config, err := config.NewAppSettingsFromFile(configFilePath)
	expectedConfig := "host=localhost port=5432 dbname=development search_path=accountkit user=deepakkv sslmode=disable"
	assert.Nil(t, err)
	assert.Equal(t, expectedConfig, config.DB.ConnectionString())
}

func TestNewAppSettingsFromFileForInvalidFile(t *testing.T) {
	configFilePath := "temp.conf"
	config, err := config.NewAppSettingsFromFile(configFilePath)
	expectedError := "Error reading config file: open temp.conf: no such file or directory"
	assert.Nil(t, config)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())
}