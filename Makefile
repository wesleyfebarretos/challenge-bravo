restart-services:
	@docker compose down --volumes
	@docker compose up -d

rebuild:
	@docker compose down --volumes
	@docker compose build
	@docker compose up -d
