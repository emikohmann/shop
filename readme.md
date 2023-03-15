# Shop

This project implements frontend and backend of an e-commerce based on microservices.

<img width="1792" alt="Screenshot 2023-03-15 at 03 53 57" src="https://user-images.githubusercontent.com/7863795/225230399-110410c2-900a-4e10-b45c-61f994654e14.png">

<img width="1176" alt="Screenshot 2023-03-15 at 03 54 25" src="https://user-images.githubusercontent.com/7863795/225230410-558f9c13-d7a6-4cda-84aa-aa6c96d8e962.png">

The backend is composed of a set of APIs:
- Item API (Golang + MongoDB + RabbitMQ + Prometheus)
- Users API (Golang + MySQL + RabbitMQ + Prometheus)
- Stores API (Golang + MySQL + RabbitMQ + Prometheus)
- Orders API (Golang + MongoDB + RabbitMQ + Prometheus)
- Discounts API (Golang + MongoDB + RabbitMQ + Prometheus)

The frontend is composed by 2 modules:
- Server (NodeJS)
- Client (NodeJS + Express + React)

## Run locally

1. Install and run MongoDB

```sh
brew install mongodb-community
brew services start mongodb-community
```

2. Install and run MySQL

```sh
brew install mysql
brew services start mysql
```

3. Install and run RabbitMQ

```sh
brew install rabbitmq
brew services start rabbitmq
rabbitmq-server
```

4. Install and run Prometheus

```shell
cd ~/prometheus-server/prometheus-2.42.0.darwin-amd64
./prometheus --config.file="prometheus.yml"
```

5. Install and run Grafana

```shell
brew install grafana
brew services start grafana
```

6. Run items-api

```shell
cd backend/items-api
go run cmd/items-api/main.go
```

7. Run users-api

```shell
cd backend/items-api
go run cmd/items-api/main.go
```

8. Run frontend server

```shell
cd frontend/server
npm run start
```

9. Run frontend client

```shell
cd frontend/client
npm run start
```

10. Update /etc/hosts

```
grep -qxF '127.0.0.1 shop.com' /etc/hosts || echo '127.0.0.1 shop.com' >> /etc/hosts
```

11. Go to http://shop.com

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
