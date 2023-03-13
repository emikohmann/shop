export DOCKER_BUILDKIT=0
export COMPOSE_DOCKER_CLI_BUILD=0
echo 'Building custom prometheus image'
cd prometheus
docker rmi prometheus:latest || true
docker build -t prometheus .
cd ..
echo 'Building items-api image'
cd backend
cd items-api
docker rmi items-api:latest
docker build -t items-api .
cd ..
cd ..
echo 'Building fe-server image'
cd frontend
cd server
docker build -t fe-server .
cd ..
cd ..
echo 'Building fe-client image'
cd frontend
cd client
docker build -t fe-client .
cd ..
cd ..
echo 'Removing previous existing images'
docker rm mongodb-dev || true
docker rm prometheus-dev || true
docker rm rabbitmq-dev || true
docker rm grafana-dev || true
docker rm items-dev || true
docker rm fe-server-dev || true
docker rm fe-client-dev || true
echo 'Running application with docker-compose'
docker-compose -p shop up -d
