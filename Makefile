.PHONY: build run lint docker-build

DOCKER_COMPOSE = docker-compose
GOLINT = golangci-lint

docker-build:
	@echo "Building and running with Docker Compose..."
	$(DOCKER_COMPOSE) up --build

run:
	@echo "Running with Docker Compose..."
	$(DOCKER_COMPOSE) up

down:
	@echo "Stopping and removing containers..."
	$(DOCKER_COMPOSE) down

lint:
	@echo "Linting the Go files..."
	$(GOLINT) run ./...

vet:
	@echo "Running go vet..."
	go vet ./...

build:
	@echo "Build server and client..."
	go build -o server cmd/server/main.go
	go build -o server cmd/client/main.go
