package config

import "github.com/caarlos0/env"

type Config struct {
	App    APP
	HTTP   HTTP
	Docker Docker
}

type APP struct {
	Name string `env:"APP_NAME" envDefault:"items-api"`
}

type HTTP struct {
	Port int `env:"HTTP_PORT" envDefault:"9999"`
}

type Docker struct {
	APIVersion string `env:"DOCKER_API_VERSION" envDefault:"1.41"`
}

// Read loads all application config
func Read() (*Config, error) {
	var config Config
	for _, target := range []interface{}{
		&config,
		&config.App,
		&config.HTTP,
		&config.Docker,
	} {
		if err := env.Parse(target); err != nil {
			return nil, err
		}
	}
	return &config, nil
}
