package config

import (
	"errors"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `env:"ENV" envDefault:"dev"`
	Port       string `env:"PORT" envDefault:"8080"`
	Database   string `env:"DATABASE_URL"`
	ProductApi string `env:"PRODUCT_API_URL"`
	MasterApi  string `env:"MASTER_API_URL"`
}

func NewConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load .env file")
	}

	cfg := new(Config)

	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config file")
	}

	return cfg, nil
}
