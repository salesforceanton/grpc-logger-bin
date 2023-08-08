package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	ENV_PREFIX = "loggerbin"
)

type Config struct {
	DBUri      string `envconfig:"DB_URI"`
	DBUsername string `envconfig:"DB_USERNAME"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_DATABASE"`
	ServerPort int    `envconfig:"SERVER_PORT"`
}

// Recieve configuration values from env variables
func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	var cfg Config
	if err := envconfig.Process(ENV_PREFIX, &cfg); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	return &cfg, nil
}
