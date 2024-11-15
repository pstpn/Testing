package config

import (
	"github.com/spf13/viper"
)

const configPath = "config/config.yaml"

type Config struct {
	Logger LoggerConfig `yaml:"logger"`
	HTTP   HTTPConfig   `yaml:"http"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type HTTPConfig struct {
	Port  int         `yaml:"port"`
	Admin AdminConfig `yaml:"admin"`
}

type AdminConfig struct {
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
}

func NewConfig() (*Config, error) {
	var err error
	var config Config

	viper.SetConfigFile(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
