# Items API

This is an example service for items that includes endpoints in order to:

- Create an item
- Get an item
- Update an item
- Delete an item

## Pre-requisites

- Install MongoDB https://www.mongodb.com/docs/manual/installation/
- Install RabbitMQ https://www.rabbitmq.com/download.html
- Install Prometheus https://prometheus.io/docs/prometheus/latest/installation/
- Install Grafana https://grafana.com/docs/grafana/latest/setup-grafana/installation/
- Install Swag https://github.com/swaggo/swag

#### Test the application with:

```bash
make test
```

#### Run the application with:

```bash
make start
make run
```

#### Stop the application with:

```bash
make stop
```

#### Generate docs with:

```bash
make docs 
```
