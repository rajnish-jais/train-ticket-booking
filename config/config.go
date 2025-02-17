package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config struct to map the config file
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

// LoadConfig reads the config.yaml file and returns a Config struct
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

// MustLoadConfig loads config and logs fatal error if loading fails
func MustLoadConfig(filename string) *Config {
	config, err := LoadConfig(filename)
	if err != nil {
		log.Fatalf("⚠️ Failed to load config: %v", err)
	}
	return config
}
