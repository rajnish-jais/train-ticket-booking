package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadConfig_Success tests if LoadConfig successfully loads a valid config file
func TestLoadConfig_Success(t *testing.T) {
	// Create a temporary valid config file
	tempFile := "test_config.yaml"
	content := "server:\n  port: \"50051\"\n"
	err := os.WriteFile(tempFile, []byte(content), 0644)
	assert.NoError(t, err)
	defer os.Remove(tempFile) // Cleanup after test

	// Load the config
	cfg, err := LoadConfig(tempFile)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "50051", cfg.Server.Port)
}

// TestLoadConfig_FileNotFound tests if LoadConfig returns an error when the file does not exist
func TestLoadConfig_FileNotFound(t *testing.T) {
	cfg, err := LoadConfig("non_existent.yaml")
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

// TestLoadConfig_InvalidYAML tests if LoadConfig returns an error when YAML is invalid
func TestLoadConfig_InvalidYAML(t *testing.T) {
	tempFile := "invalid_config.yaml"
	content := "server:\n  port: 50051\n invalid_yaml"
	err := os.WriteFile(tempFile, []byte(content), 0644)
	assert.NoError(t, err)
	defer os.Remove(tempFile) // Cleanup

	cfg, err := LoadConfig(tempFile)
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

// TestMustLoadConfig_Success ensures MustLoadConfig loads a valid config file
func TestMustLoadConfig_Success(t *testing.T) {
	tempFile := "test_config.yaml"
	content := "server:\n  port: \"50051\"\n"
	err := os.WriteFile(tempFile, []byte(content), 0644)
	assert.NoError(t, err)
	defer os.Remove(tempFile)

	cfg := MustLoadConfig(tempFile)
	assert.NotNil(t, cfg)
	assert.Equal(t, "50051", cfg.Server.Port)
}
