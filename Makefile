COMPOSE_FILE=docker-compose.dev.yaml


dev:
	@docker-compose -f $(COMPOSE_FILE) up -d --remove-orphans

down:
	@docker-compose -f $(COMPOSE_FILE) down

stop:
	@docker-compose -f $(COMPOSE_FILE) stop

purge:
	@docker-compose -f $(COMPOSE_FILE) down -v