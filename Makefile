include ./app/.env

DATABASE=${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable

start-services:
	@docker compose up -d

restart-services:
	@docker compose down --volumes
	@docker compose up -d

restart-db:
	@docker compose down postgres --volumes
	@docker compose up postgres -d

rebuild:
	@docker compose down --volumes
	@docker compose build
	@docker compose up -d