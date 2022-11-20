package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	Timeout  int `env:"TIMEOUT" default:"60"`
	Postgres Postgres
}

type Postgres struct {
	URL      string `env:"DATABASE_URL" default:"postgres://postgres:postgres@localhost:5432/postgres"`
	Host     string `env:"POSTGRES_HOST" default:"localhost"`
	Port     int    `env:"POSTGRES_PORT" default:"5432"`
	User     string `env:"POSTGRES_USER" default:"postgres"`
	Password string `env:"POSTGRES_PASS" default:"postgres"`
	Database string `env:"POSTGRES_DATABASE" default:"postgres"`
	Timeout  int    `env:"POSTGRES_TIMEOUT" default:"30"`
}

func New(filenames ...string) (*Config, error) {
	cfg := new(Config)

	if len(filenames) > 0 {
		if err := godotenv.Load(filenames...); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	defaults.SetDefaults(cfg)

	return cfg, nil
}
