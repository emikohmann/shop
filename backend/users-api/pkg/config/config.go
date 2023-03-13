package config

import "github.com/caarlos0/env"

type Config struct {
    App           APP
    HTTP          HTTP
    UsersMySQL    UsersMySQL
    UsersRabbitMQ UsersRabbitMQ
}

type APP struct {
    Name string `env:"APP_NAME" envDefault:"users-api"`
}

type HTTP struct {
    Port int `env:"HTTP_PORT" envDefault:"8081"`
}

type UsersMySQL struct {
    Host     string `env:"USERS_MYSQL_DB_HOST" envDefault:"mysql-dev"`
    Port     int    `env:"USERS_MYSQL_DB_PORT" envDefault:"3306"`
    Database string `env:"USERS_MYSQL_DB_DATABASE" envDefault:"users"`
    User     string `env:"USERS_MYSQL_USERNAME" envDefault:"root"`
    Password string `env:"USERS_MYSQL_PASSWORD" envDefault:""`
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
        &config.UsersMySQL,
        &config.UsersRabbitMQ,
    } {
        if err := env.Parse(target); err != nil {
            return nil, err
        }
    }
    return &config, nil
}
