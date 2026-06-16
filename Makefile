.PHONY: run worker build test lint migrate-up migrate-down docker-up docker-down

run:
	go run ./cmd/api

worker:
	go run ./cmd/worker

build:
	go build -ldflags="-s -w" -o bin/api ./cmd/api

test:
	go test -race ./...

lint:
	go vet ./...

migrate-up:
	@echo "Configure golang-migrate, then: migrate -path migrations -database \$$DATABASE_URL up"

migrate-down:
	@echo "Configure golang-migrate, then: migrate -path migrations -database \$$DATABASE_URL down 1"

docker-up:
	docker compose up -d

docker-down:
	docker compose down
