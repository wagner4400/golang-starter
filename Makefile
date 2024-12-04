# Simple Makefile for a Go project

# Build the application
all: build

swagger:
	go install github.com/swaggo/swag/cmd/swag@v1.16.3
	swag init \
		--parseInternal \
		--parseDependency \
		--parseDepth 3 \
		--output docs \
		--generalInfo ./cmd/app/main.go

test:
	go clean -testcache
	go test ./... | grep -v '?'

# Run the application
run:
	go run cmd/api/main.go

tidy:
	go mod tidy
	go mod vendor

build:
	@echo "Building..."
	
	
	@go build -o main.exe cmd/api/main.go


# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi


# Test the application
test:
	@echo "Testing..."
	@go test ./... -v


# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v


# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload

watch:
	@air


.PHONY: all build run test clean watch
