package config

import "github.com/caarlos0/env"

type Config struct {
    App           APP
    HTTP          HTTP
    UsersMongoDB  UsersMongoDB
    UsersRabbitMQ UsersRabbitMQ
}

type APP struct {
    Name string `env:"APP_NAME" envDefault:"users-api"`
}

type HTTP struct {
    Port int `env:"HTTP_PORT" envDefault:"8081"`
}

type UsersMongoDB struct {
    Host       string `env:"USERS_MONGO_DB_HOST" envDefault:"mongodb-dev"`
    Port       int    `env:"USERS_MONGO_DB_PORT" envDefault:"27017"`
    Database   string `env:"USERS_MONGO_DB_DATABASE" envDefault:"users-api"`
    Collection string `env:"USERS_MONGO_DB_COLLECTION" envDefault:"users"`
}

type UsersRabbitMQ struct {
    Host      string `env:"USERS_RABBIT_MQ_HOST" envDefault:"rabbitmq-dev"`
    Port      int    `env:"USERS_RABBIT_MQ_PORT" envDefault:"5672"`
    User      string `env:"USERS_RABBIT_MQ_USER" envDefault:"guest"`
    Password  string `env:"USERS_RABBIT_MQ_PASSWORD" envDefault:"guest"`
    QueueName string `env:"USERS_RABBIT_MQ_QUEUE_NAME" envDefault:"users"`
}

// Read loads all application config
func Read() (*Config, error) {
    var config Config
    for _, target := range []interface{}{
        &config,
        &config.App,
        &config.HTTP,
        &config.UsersMongoDB,
        &config.UsersRabbitMQ,
    } {
        if err := env.Parse(target); err != nil {
            return nil, err
        }
    }
    return &config, nil
}
