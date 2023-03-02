# Items API

This is an example service for items.

It basically includes 4 endpoint in order to:
- Collect metrics

- Create an item
- Get an item
- Update an item
- Delete an item

And it works with the following services:
- MongoDB as the source of truth
- RabbitMQ as the message queue for notifications
- Prometheus as the metrics collector to visualize in Grafana

## Test

Test the application with:

```bash
go test ./... -v
```

## Pre-requisites

- Install MongoDB
- Install RabbitMQ
- Install Prometheus
- Install Grafana

## Run

Run the application with:

```bash
brew services start mongodb-community
```

```bash
rabbitmq-server
```

```bash
./prometheus --config.file=prometheus.yml
```

```bash
brew services start grafana
```

```bash
cd backend/items-api/cmd/items-api
go build
./items-api
```