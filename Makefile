# Variables
PROJECT_NAME := sso
BINARY_NAME := $(PROJECT_NAME)
GO := go
BUILD_DIR := ./bin
SRC_DIR := ./cmd/$(PROJECT_NAME)

all: help

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)
	@echo "Binary built: $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	@echo "Cleaned."

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Available targets:"
	@echo "  build       Compile the Go application"
	@echo "  clean       Remove build artifacts"
	@echo "  help        Show this help message"

.PHONY: all build clean help