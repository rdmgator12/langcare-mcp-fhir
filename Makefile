.PHONY: build test test-coverage lint clean run run-http install-tools

# Build the server binary
build:
	@echo "Building langcare-mcp-fhir server..."
	@go build -o bin/langcare-mcp-fhir ./cmd/server

# Run all tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/ coverage.out coverage.html

# Run server in stdio mode
run:
	@echo "Starting server (stdio mode)..."
	@go run cmd/server/main.go

# Run server in HTTP mode
run-http:
	@echo "Starting server (HTTP mode)..."
	@go run cmd/server/main.go --transport=http --port=8080

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
