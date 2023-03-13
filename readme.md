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

## Requirements and execution

In order to run the project, only Docker is needed, since all services are configured and executed with docker-compose.

1. Install Docker
2. Run:

```sh
source run.sh
```

Happy coding! :)