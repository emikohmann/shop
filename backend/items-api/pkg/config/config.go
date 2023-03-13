package config

import "github.com/caarlos0/env"

type Config struct {
	App           APP
	HTTP          HTTP
	ItemsMongoDB  ItemsMongoDB
	ItemsRabbitMQ ItemsRabbitMQ
}

type APP struct {
	Name string `env:"APP_NAME" envDefault:"items-api"`
}

type HTTP struct {
	Port int `env:"HTTP_PORT" envDefault:"8080"`
}

type ItemsMongoDB struct {
	Host       string `env:"ITEMS_MONGO_DB_HOST" envDefault:"mongodb-dev"`
	Port       int    `env:"ITEMS_MONGO_DB_PORT" envDefault:"27017"`
	Database   string `env:"ITEMS_MONGO_DB_DATABASE" envDefault:"items-api"`
	Collection string `env:"ITEMS_MONGO_DB_COLLECTION" envDefault:"items"`
}

type ItemsRabbitMQ struct {
	Host      string `env:"ITEMS_RABBIT_MQ_HOST" envDefault:"rabbitmq-dev"`
	Port      int    `env:"ITEMS_RABBIT_MQ_PORT" envDefault:"5672"`
	User      string `env:"ITEMS_RABBIT_MQ_USER" envDefault:"guest"`
	Password  string `env:"ITEMS_RABBIT_MQ_PASSWORD" envDefault:"guest"`
	QueueName string `env:"ITEMS_RABBIT_MQ_QUEUE_NAME" envDefault:"items"`
}

// Read loads all application config
func Read() (*Config, error) {
	var config Config
	for _, target := range []interface{}{
		&config,
		&config.App,
		&config.HTTP,
		&config.ItemsMongoDB,
		&config.ItemsRabbitMQ,
	} {
		if err := env.Parse(target); err != nil {
			return nil, err
		}
	}
	return &config, nil
}
