package config

import "github.com/caarlos0/env"

type Config struct {
	HTTP         HTTP
	ItemsMongoDB ItemsMongoDB
}

type HTTP struct {
	Port int `env:"HTTP_PORT" envDefault:"8080"`
}

type ItemsMongoDB struct {
	Host       string `env:"ITEMS_MONGO_DB_HOST" envDefault:"localhost"`
	Port       int    `env:"ITEMS_MONGO_DB_PORT" envDefault:"27017"`
	Database   string `env:"ITEMS_MONGO_DB_DATABASE" envDefault:"items-api"`
	Collection string `env:"ITEMS_MONGO_DB_COLLECTION" envDefault:"items"`
}

// Read loads all application config
func Read() (*Config, error) {
	var config Config
	for _, target := range []interface{}{
		&config,
		&config.HTTP,
		&config.ItemsMongoDB,
	} {
		if err := env.Parse(target); err != nil {
			return nil, err
		}
	}
	return &config, nil
}
