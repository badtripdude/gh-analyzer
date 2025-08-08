DOCKER_COMPOSE := docker compose

docker-up:
	$(DOCKER_COMPOSE) up --build -d

docker-clear:
	$(DOCKER_COMPOSE) down -v

docker-down:
	$(DOCKER_COMPOSE) down

# kafka tasks for convenient testing
kafka-up:
	$(DOCKER_COMPOSE) -f ../kafka/docker-compose.kafka.yml up -d

kafka-down:
	$(DOCKER_COMPOSE) -f ../kafka/docker-compose.kafka.yml down -v

logs:
	$(DOCKER_COMPOSE) logs -f app