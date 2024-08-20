.PHONY: all build test swagger golangci-lint help

# Go parameters
GOCMD=go1.20
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLANGCI_LINT=golangci-lint run
BINARY_NAME=ai-audio-service
BUILD_DIR=build

all: build

build: 
	@echo "Building the application..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) cmd/main.go

test: 
	@echo "Running tests..."
	$(GOTEST) ./...

swagger:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/main.go -o docs

golangci-lint:
	@echo "Running golangci-lint..."
	$(GOLANGCI_LINT)

help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  test          - Run tests"
	@echo "  swagger       - Generate Swagger documentation"
	@echo "  golangci-lint - Run golangci-lint"
	@echo "  help          - Show this help message"