.PHONY: build build-all build-go test test-coverage lint clean run run-http install-tools deps apps apps-clean

# Build MCP App HTML bundles and copy to Go embed directory
apps:
	@bash scripts/build-apps.sh

# Clean app build artifacts
apps-clean:
	@rm -rf apps/dist apps/dist-tmp apps/node_modules
	@find internal/apps/dist -name '*.html' ! -name '.gitkeep.html' -delete 2>/dev/null || true

# Build the server binary (includes embedded apps)
build: apps
	@echo "Building langcare-mcp-fhir server..."
	@go build -o bin/langcare-mcp-fhir ./cmd/server

# Build Go binary only (skip apps rebuild, for Go-only changes)
build-go:
	@echo "Building langcare-mcp-fhir server (Go only)..."
	@go build -o bin/langcare-mcp-fhir ./cmd/server

# Build for all platforms (for npm publishing)
build-all: clean apps
	@echo "Building for all platforms..."
	@mkdir -p bin
	@echo "Building for macOS (Intel)..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/langcare-mcp-fhir-darwin-amd64 ./cmd/server
	@echo "Building for macOS (ARM64)..."
	@GOOS=darwin GOARCH=arm64 go build -o bin/langcare-mcp-fhir-darwin-arm64 ./cmd/server
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o bin/langcare-mcp-fhir-linux-amd64 ./cmd/server
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/langcare-mcp-fhir-windows-amd64.exe ./cmd/server
	@echo "✓ All platforms built successfully"

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
	@go run cmd/server/main.go -http -port 8080

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy