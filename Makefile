.PHONY: help build run dev test swagger swagger-fmt

# Default target
.DEFAULT_GOAL := help

## help: Show available make commands
help:
	@echo "Usage:"
	@echo "  make swagger       - Regenerate Swagger API documentation"
	@echo "  make dev           - Run server with live reload (using air or go run)"
	@echo "  make build         - Build the server binary"
	@echo "  make test          - Run all unit tests"
	@echo "  make fmt           - Format code using gofmt"
	@echo "  make vet           - Run static code analysis with go vet"

## swagger: Regenerate Swagger documentation in docs/
swagger:
	@echo "Generating Swagger documentation..."
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go -o docs

## dev: Run application with live-reload (Air or fallback to go run)
dev: swagger
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air is not installed, running with 'go run ./cmd/server'..."; \
		go run ./cmd/server; \
	fi

## build: Compile server binary into bin/server
build: swagger
	@echo "Building server binary..."
	@mkdir -p bin
	go build -o bin/server ./cmd/server

## test: Run unit tests across all packages
test:
	@echo "Running tests..."
	go test -v ./...

## fmt: Format all Go files
fmt:
	@echo "Formatting Go files..."
	gofmt -s -w .

## vet: Run go vet static analysis
vet:
	@echo "Running go vet..."
	go vet ./...
