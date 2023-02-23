package config

import "github.com/caarlos0/env"

type Config struct {
	HTTP HTTP
}

type HTTP struct {
	Port int `env:"HTTP_PORT" envDefault:"8080"`
}

// Read loads all application config
func Read() (*Config, error) {
	var config Config
	for _, target := range []interface{}{
		&config,
		&config.HTTP,
	} {
		if err := env.Parse(target); err != nil {
			return nil, err
		}
	}
	return &config, nil
}
