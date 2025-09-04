package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	GRPCServerConfig struct {
		Port string `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"order-service"`

	DatabaseConfig struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"order-database"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read config file: %w", err)
	}

	expanded := os.ExpandEnv(string(data))

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(strings.NewReader(expanded)); err != nil {
		return config, fmt.Errorf("viper read config: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("unmarshal config: %w", err)
	}

	return config, nil
}
