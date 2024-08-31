include .env

DATABASE=${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable

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

# Migrations
create-table:
	@migrate create -ext=sql -dir=./internal/migration -seq $(shell echo $(filter-out $@,$(MAKECMDGOALS))_table)

create-seed:
	@migrate create -ext sql -dir ./internal/migration -seq $(shell echo $(filter-out $@,$(MAKECMDGOALS))_seeder)

create-view:
	@migrate create -ext sql -dir ./internal/migration -seq $(shell echo $(filter-out $@,$(MAKECMDGOALS))_view)

create-schema:
	@migrate create -ext sql -dir ./internal/migration -seq $(shell echo $(filter-out $@,$(MAKECMDGOALS))_schema)

migrations-up:
	@migrate -path ./internal/migration -database $(DATABASE) -verbose up

migrations-down:
	@migrate -path ./internal/migration -database $(DATABASE) -verbose down -all