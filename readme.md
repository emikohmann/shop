# Shop

This project implements the backend and the frontend of an e-commerce using microservices.

The backend is composed of a set of APIs:
- Item API (built with Golang)
- Users API (built with Golang)
- Stores API (built with Golang)
- Orders API (built with Golang)
- Discounts API (built with Golang)

And it uses a set of services in order to run:
- MongoDB
- RabbitMQ
- Prometheus
- Grafana
- MySQL

The frontend is composed by 2 modules:
- Server (built with Node)
- Client (built with Node, Express and React)

## Run locally

1. Install and run MongoDB
```sh
brew install mongodb-community
brew services start mongodb-community
```

2. Install and run RabbitMQ
```sh
brew install rabbitmq
brew services start rabbitmq
rabbitmq-server
```

3. Install and run Prometheus
```shell
cd ~/prometheus-server/prometheus-2.42.0.darwin-amd64
./prometheus --config.file="prometheus.yml"
```

4. Install and run Grafana
```shell
brew install grafana
brew services start grafana
```

5. Run items-api
```shell
cd backend/items-api
go run cmd/items-api/main.go
```

7. Run frontend server
```shell
cd frontend/server
npm run start
```

8. Run frontend client
```shell
cd frontend/client
npm run start
```

9. Update /etc/hosts
```
grep -qxF '127.0.0.1 shop.com' /etc/hosts || echo '127.0.0.1 shop.com' >> /etc/hosts
```

10. Go to http://shop.com

## Run with docker

1. Install Docker

2. Generate custom images
```
cd backend/{image-name}
docker rmi {image-name}:latest
docker build -t {image-name} .
```

3. Run
```shell
docker-compose -p shop up -d
```

4. Update /etc/hosts
```
grep -qxF '127.0.0.1 shop.com' /etc/hosts || echo '127.0.0.1 shop.com' >> /etc/hosts
```

5. Go to http://shop.com

## Generate and expose swagger (backend services)

```shell
cd backend/{api-name}
swag fmt && swag init -g cmd/{api-name}/main.go --output docs/openapi
```

## Generate mocks (backend services)

```shell
mockery --dir=internal/application --name={interfaceName} --output=internal/mocks/{pkgName} --filename={mockName}_mock.go --outpkg={pkgName} --exported
```

Happy coding! :)