start-services:
	@echo "Restarting services"
	brew services stop mongodb-community
	brew services start mongodb-community
	brew services stop rabbitmq
	brew services start rabbitmq
	brew services stop grafana
	brew services start grafana
	@echo "Starting servers"
	rabbitmq-server & cd ~/prometheus-server/prometheus-2.42.0.darwin-amd64 && ./prometheus --config.file="prometheus.yml"

run-application:
	go run cmd/items-api/main.go

stop-services:
	launchctl unload ~/Library/LaunchAgents/homebrew.mxcl.rabbitmq.plist
	brew services stop mongodb-community
	brew services stop rabbitmq
	brew services stop grafana

test-application:
	go test ./... -v

generate-docs:
	go install github.com/swaggo/swag/cmd/swag
	swag fmt && swag init -g cmd/items-api/main.go --output docs/openapi