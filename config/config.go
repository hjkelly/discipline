package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	Port       int    `env:"PORT"`
	APIBaseURL string `env:"API_BASE_URL"`
	SentryDSN  string `env:"SENTRY_DSN"`
}

func ParseConfig() (*Config, error) {
	configPtr := new(Config)
	err := env.Parse(configPtr)
	if err != nil {
		return nil, err
	}

	if configPtr.Port == 0 {
		return nil, fmt.Errorf("You must provide the PORT environment variable.")
	} else if configPtr.SentryDSN == "" {
		return nil, fmt.Errorf("You must provide the API_BASE_URL environment variable.")
	} else if configPtr.SentryDSN == "" {
		return nil, fmt.Errorf("You must provide the SENTRY_DSN environment variable.")
	}

	return configPtr, nil
}
