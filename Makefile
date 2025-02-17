# Variables
APP_NAME = train-ticket-app
DOCKER_IMAGE = train-ticket-app
SERVER_BINARY = train-ticket-server
CLIENT_BINARY = train-ticket-client
PROTO_SRC = proto/train_ticket.proto
PROTO_OUT = .

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod tidy

# Generate gRPC files from proto definitions
proto:
	@echo "Generating gRPC code..."
	protoc --proto_path=proto --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) $(PROTO_SRC)

# Build the server binary
build-server:
	@echo "Building server binary..."
	go build -o $(SERVER_BINARY) cmd/server/main.go

# Build the client binary
build-client:
	@echo "Building client binary..."
	go build -o $(CLIENT_BINARY) cmd/client/main.go

# Run the gRPC server
run-server: build-server
	@echo "Starting the gRPC server..."
	./$(SERVER_BINARY)

# Run the gRPC client
run-client: build-client
	@echo "Starting the gRPC client..."
	./$(CLIENT_BINARY)

# Run all tests
test:
	@echo "Running tests..."
	go test ./... -v

coverage:
	@echo "Running tests with coverage..."
	go test ./... -coverprofile=coverage.out -v -timeout=60s
	@echo "Generating HTML coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Opening coverage report..."
	@if command -v xdg-open >/dev/null; then \
		xdg-open coverage.html; \
	elif command -v open >/dev/null; then \
		open coverage.html; \
	elif command -v start >/dev/null; then \
		start coverage.html; \
	else \
		echo "Please open coverage.html manually"; \
	fi


# Clean up binaries and generated files
clean:
	@echo "Cleaning up..."
	rm -f $(SERVER_BINARY) $(CLIENT_BINARY) coverage.out coverage.html
	rm -rf $(PROTO_OUT)/proto/*.pb.go


# Default target
all: install proto build-server build-client

.PHONY: install proto build-server build-client run-server run-client test coverage clean all
