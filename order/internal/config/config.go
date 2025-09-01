package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPCServerConfig struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"order-service"`

	DatabaseConfig struct {
		Url string `yaml:"url"`
	} `yaml:"order-database"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
