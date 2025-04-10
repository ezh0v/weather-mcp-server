package server

import (
	"errors"
	"time"
)

type Config struct {
	ListenAddr        string
	WeatherAPIKey     string
	WeatherAPITimeout time.Duration
}

func (c *Config) Validate() error {
	if c.WeatherAPIKey != "" {
		return nil
	}

	return errors.New("WeatherAPIKey is required")
}
