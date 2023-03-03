# Items API

This is an example service for items that includes endpoints in order to:

- Create an item
- Get an item
- Update an item
- Delete an item

## Pre-requisites

- Install and run MongoDB https://www.mongodb.com/docs/manual/installation/
- Install and run RabbitMQ https://www.rabbitmq.com/download.html
- Install and run Prometheus https://prometheus.io/docs/prometheus/latest/installation/
- Install and run Grafana https://grafana.com/docs/grafana/latest/setup-grafana/installation/
- Install and run Swag https://github.com/swaggo/swag

#### Test the application with:

```bash
go test ./... -v
```

#### Run the application with:

```bash
go run cmd/items-api/main.go
```

#### Generate docs with:

```bash
swag fmt && swag init -g cmd/items-api/main.go --output docs/openapi 
```
