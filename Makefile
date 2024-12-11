# Variables
APP_NAME=books
APP_PORT=8080
DB_FILE=books.db

# Default target
.DEFAULT_GOAL := help

.PHONY: help run build clean test

# Show help
help:
	@echo "Available commands:"
	@echo "  make run        - Run the application"
	@echo "  make build      - Build the application binary"
	@echo "  make clean      - Clean up the binary and database file"
	@echo "  make test       - Run tests"

# Run the application
run:
	@echo "Running the application on http://localhost:$(APP_PORT)..."
	@go run ./cmd/main/main.go

# Build the application binary
build:
	@echo "Building the binary..."
	@go build -o $(APP_NAME)

# Clean up generated files
clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME) $(DB_FILE)

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

