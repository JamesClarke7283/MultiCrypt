# Define variables
BUILD_DIR = build
SRC_DIR = src
BENCH_DIR = bench
TARGET = $(BUILD_DIR)/multicrypt
PROFILE_FILE = $(BUILD_DIR)/default.pgo
PGO_TARGET = $(BUILD_DIR)/multicrypt_pgo

# Default target
all: build

# Generate profile data
profile:
	mkdir -p $(BUILD_DIR)
	go build -o $(TARGET).profile $(SRC_DIR)/main.go
	$(TARGET).profile -cpuprofile=$(PROFILE_FILE) &
	sleep 10 # Run your application for a representative amount of time
	pkill -f $(TARGET).profile

# Build the application without PGO
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(TARGET) $(SRC_DIR)/main.go

# Build the application with PGO
build-with-pgo:
	@if [ ! -f $(PROFILE_FILE) ]; then \
		echo "Profile file not found. Please run 'make profile' first."; \
		exit 1; \
	fi
	mkdir -p $(BUILD_DIR)
	go build -o $(PGO_TARGET) -pgo=$(PROFILE_FILE) $(SRC_DIR)/main.go

# Clean up build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Run the application
run: build
	$(TARGET)

# Benchmark PGO performance
benchmark: build build-with-pgo
	go run $(BENCH_DIR)/pgo_benchmark.go $(TARGET) $(PGO_TARGET)

.PHONY: all build build-with-pgo clean run profile benchmark