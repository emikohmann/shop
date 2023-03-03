start:
	@echo "Restarting services and starting servers"
	brew services stop mongodb-community
	brew services start mongodb-community
	brew services stop rabbitmq
	brew services start rabbitmq
	brew services stop grafana
	brew services start grafana
	rabbitmq-server & cd ~/prometheus-server/prometheus-2.42.0.darwin-amd64 && ./prometheus --config.file="prometheus.yml"

run:
	@echo "Running application"
	go run cmd/items-api/main.go

stop:
	@echo "Stopping application"
	brew services stop mongodb-community
	brew services stop rabbitmq
	brew services stop grafana

test:
	@echo "Testing application"
	go test ./... -v

docs:
	@echo "Generating docs"
	swag fmt && swag init -g cmd/items-api/main.go --output docs/openapi