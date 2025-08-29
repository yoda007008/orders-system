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
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"dbname"`
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
