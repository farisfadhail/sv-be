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