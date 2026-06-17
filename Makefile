.PHONY: run build test lint migrate-up migrate-down docker-up docker-down clean

# Run the API server locally
run:
	go run ./cmd/api

# Run the background worker
worker:
	go run ./cmd/worker

# Build production binary
build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/api ./cmd/api

# Run all tests with race detector
test:
	go test -race ./...

# Run static analysis
lint:
	staticcheck ./...
	go vet ./...

# Apply database migrations
migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

# Rollback one migration
migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1

# Start Docker services
docker-up:
	docker-compose up -d

# Stop Docker services
docker-down:
	docker-compose down

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Format Go code
fmt:
	go fmt ./...

# Download dependencies
deps:
	go mod download
	go mod tidy