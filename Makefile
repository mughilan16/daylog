APP_NAME := daylog 
BUILD_DIR := bin

.PHONY: all build run clean fmt

all: build

build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME)

run:
	@go run .

fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

