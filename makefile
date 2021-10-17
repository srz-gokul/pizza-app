## Build and start all defined containers
up:
	docker-compose  up --build

## Stop and remove all defined containers
down:
	docker-compose down

## View logs
log:
	docker-compose  logs --tail 100 api