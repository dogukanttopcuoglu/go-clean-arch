package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	HTTP http
	JWT  jwt
	PG   pg
}

type http struct {
	Port string `env:"HTTP_PORT,required"`
}

type pg struct {
	PoolMax int    `env:"PG_POOL_MAX,required"`
	URL     string `env:"PG_URL,required"`
}

type jwt struct {
	Secret      string        `env:"JWT_SECRET,required"`
	TokenExpiry time.Duration `env:"JWT_TOKEN_EXPIRY" envDefault:"24h"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
