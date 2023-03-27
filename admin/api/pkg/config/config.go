package config

import "github.com/caarlos0/env"

type Config struct {
	App  APP
	HTTP HTTP
}

type APP struct {
	Name string `env:"APP_NAME" envDefault:"items-api"`
}

type HTTP struct {
	Port int `env:"HTTP_PORT" envDefault:"9999"`
}

// Read loads all application config
func Read() (*Config, error) {
	var config Config
	for _, target := range []interface{}{
		&config,
		&config.App,
		&config.HTTP,
	} {
		if err := env.Parse(target); err != nil {
			return nil, err
		}
	}
	return &config, nil
}
