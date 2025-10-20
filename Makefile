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

doc-migrate-up:
	docker-compose exec app migrate -path=./database/migrations -database "postgres://postgres:password@postgres:5432/test-sv?sslmode=disable" up

doc-migrate-down:
	docker-compose exec app migrate -path=./database/migrations -database "postgres://postgres:password@postgres:5432/test-sv?sslmode=disable" down

doc-migrate-drop:
	docker-compose exec app migrate -path=./database/migrations -database "postgres://postgres:password@postgres:5432/test-sv?sslmode=disable" drop

doc-migrate-force:
	docker-compose exec app migrate -path=./database/migrations -database "postgres://postgres:password@postgres:5432/test-sv?sslmode=disable" force $(VERSION)