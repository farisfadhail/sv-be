run:
	go run main.go

wire:
	cd internal/injector && wire

migrate-up:
	go run cmd/migration.go migrate:up

migrate-fresh:
	go run cmd/migration.go migrate:fresh

migrate-create:
	migrate create -ext sql -dir ./database/migrations $(name)

.PHONY: help up down build logs migrate-up migrate-down migrate-fresh seed clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up: ## Start all services
	docker-compose up -d

down: ## Stop all services
	docker-compose down

build: ## Build all services
	docker-compose build

rebuild: ## Rebuild and start all services
	docker-compose up -d --build

logs: ## Show logs
	docker-compose logs -f

logs-app: ## Show app logs only
	docker-compose logs -f app

logs-mysql: ## Show MySQL logs only
	docker-compose logs -f mysql

doc-migrate-up: ## Run database migrations up
	docker-compose exec app migrate -path=./database/migrations -database "mysql://root:password@tcp(mysql:3306)/test-sv" up

doc-migrate-down: ## Run database migrations down
	docker-compose exec app migrate -path=./database/migrations -database "mysql://root:password@tcp(mysql:3306)/test-sv" down

doc-migrate-drop: ## Drop all migrations
	docker-compose exec app migrate -path=./database/migrations -database "mysql://root:password@tcp(mysql:3306)/test-sv" drop

doc-migrate-force: ## Force migration version (use: make migrate-force VERSION=1)
	docker-compose exec app migrate -path=./database/migrations -database "mysql://root:password@tcp(mysql:3306)/test-sv" force $(VERSION)

seed: ## Run database seeders
	docker-compose exec app ./seeder

shell-app: ## Access app container shell
	docker-compose exec app sh

shell-mysql: ## Access MySQL shell
	docker-compose exec mysql mysql -uroot -ppassword test-sv

clean: ## Clean up containers, volumes, and images
	docker-compose down -v
	docker system prune -f

restart: down up ## Restart all services

ps: ## Show running containers
	docker-compose ps