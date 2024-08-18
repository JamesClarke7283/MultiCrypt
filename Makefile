# Define variables
BUILD_DIR = build
SRC_DIR = src
TARGET = $(BUILD_DIR)/multicrypt

# Default target
all: build

# Build the application
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(TARGET) $(SRC_DIR)/main.go

# Clean up build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Run the application
run: build
	$(TARGET)

.PHONY: all build clean run
