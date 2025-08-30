# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=littlecode-backend
MAIN_PATH=./cmd/main.go

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

# Run the application
run:
	$(GOCMD) run $(MAIN_PATH)

# Clean build files
clean:
	$(GOCLEAN)
	del $(BINARY_NAME).exe

# Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Run tests
test:
	$(GOTEST) -v ./...

# Install dependencies
install:
	$(GOGET) ./...

# Hot reload for development (requires air)
dev:
	air

.PHONY: build run clean deps test install dev
