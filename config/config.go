package config

import (
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

const configPath = "config/config.yaml"

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Auth     AuthConfig     `yaml:"auth"`
	HTTP     HTTPConfig     `yaml:"http"`
	Database DatabaseConfig `yaml:"database"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type AuthConfig struct {
	SigningKey      string        `yaml:"signingKey"`
	AccessTokenTTL  time.Duration `yaml:"accessTokenTTL"`
	RefreshTokenTTL time.Duration `yaml:"refreshTokenTTL"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewConfig() (*Config, error) {
	var err error
	var config Config
	backendType := os.Getenv("BACKEND_TYPE")

	viper.SetConfigFile(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	if backendType == "docker" {
		backendPort, err := strconv.Atoi(os.Getenv("BACKEND_PORT"))
		if err != nil {
			return nil, err
		}
		postgresPort, err := strconv.Atoi(os.Getenv("POSTGRESQL_PORT"))
		if err != nil {
			return nil, err
		}

		config.HTTP = HTTPConfig{Port: backendPort}

		config.Database.Postgres.Host = os.Getenv("POSTGRESQL_HOST")
		config.Database.Postgres.Port = postgresPort
		config.Database.Postgres.User = os.Getenv("POSTGRESQL_USERNAME")
		config.Database.Postgres.Password = os.Getenv("POSTGRESQL_PASSWORD")
		config.Database.Postgres.Database = os.Getenv("POSTGRESQL_DATABASE")
	}

	return &config, nil
}
