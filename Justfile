# Morphir Go - Build Orchestration with Just

# Default recipe
default:
    @just --list

# Build the CLI application
build:
    @echo "Building morphir CLI..."
    go build -o bin/morphir ./cmd/morphir

# Build the development version of the CLI (morphir-dev)
build-dev:
    @echo "Building morphir-dev CLI..."
    go build -o bin/morphir-dev ./cmd/morphir

# Run tests across all modules
test:
    @echo "Running tests..."
    go test ./...

# Format all Go code
fmt:
    @echo "Formatting Go code..."
    go fmt ./...

# Run linters (requires golangci-lint)
lint:
    @echo "Running linters..."
    @if command -v golangci-lint > /dev/null; then \
        golangci-lint run ./...; \
    else \
        echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
    fi

# Clean build artifacts
clean:
    @echo "Cleaning build artifacts..."
    rm -rf bin/
    go clean ./...

# Download dependencies for all modules
deps:
    @echo "Downloading dependencies..."
    go work sync
    go mod download ./...

# Run go mod tidy for all modules
mod-tidy:
    @echo "Running go mod tidy..."
    cd cmd/morphir && go mod tidy
    cd pkg/models && go mod tidy
    cd pkg/tooling && go mod tidy
    cd pkg/sdk && go mod tidy
    cd pkg/pipeline && go mod tidy

# Install the CLI using go install (installs to $GOPATH/bin or $GOBIN)
install:
    @echo "Installing morphir CLI..."
    go install ./cmd/morphir
    @echo "Installed successfully!"

# Install the development version as morphir-dev
install-dev: build-dev
    @echo "Installing morphir-dev CLI..."
    @python3 -c "import os, subprocess, shutil; gopath = subprocess.check_output(['go', 'env', 'GOPATH']).decode().strip(); gobin = subprocess.check_output(['go', 'env', 'GOBIN']).decode().strip(); target = (gobin if gobin else f'{gopath}/bin'); os.makedirs(target, exist_ok=True); shutil.copy('bin/morphir-dev', f'{target}/morphir-dev'); print(f'Installed to {target}/morphir-dev')"

# Run the CLI application
run: build
    @./bin/morphir

# Run the development version of the CLI
run-dev: build-dev
    @./bin/morphir-dev

# Verify all modules build successfully
verify:
    @echo "Verifying all modules build..."
    go build ./cmd/morphir
    go build ./pkg/models
    go build ./pkg/tooling
    go build ./pkg/sdk
    go build ./pkg/pipeline
    @echo "All modules build successfully!"
