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

## Available endpoints

### Create an item

> **POST** `/items`

```json
{
  "id": 1,
  "name": "Iphone 13 128GB 4GB RAM",
  "description": "The iPhone 13 display has rounded corners that follow a beautiful curved design, and these corners are within a standard rectangle. When measured as a standard rectangular shape, the screen is 6.06 inches diagonally (actual viewable area is less). Both models: HDR display.",
  "thumbnail": "https://contactcenter.macstation.com.ar/web/image/product.template/8551/image_256/%5BMLV93LE-A%5D%20iPhone%2013%20Pro%20128GB%20-%20Grafito?unique=ed3cc51",
  "images": [
    "https://www.macstation.com.ar/img/productos/2599-2.jpg"
  ],
  "is_active": true,
  "restrictions": [],
  "price": 729.99,
  "stock": 1
}
```

### Get an item

> **GET** `/items/1`

### Update an item

> **PUT** `/items/1`

```json
{
  "name": "Iphone 13 128GB 4GB RAM Updated"
}
```

### Delete an item

> **DELETE** `/items/1`