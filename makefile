.PHONY: all build clean run test fmt vet lint help

all: build

build:
	go build -o bin/app ./src/cmd/main.go

clean:
	rm -rf bin/

run: build
	./bin/app

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

help:
	@echo "Makefile options:"
	@echo "  all    - Default target, builds the application"
	@echo "  build  - Compiles the Go application"
	@echo "  clean  - Removes the bin/ directory"
	@echo "  run    - Builds and runs the application"
	@echo "  test   - Runs all tests"
	@echo "  fmt    - Formats the Go code"
	@echo "  vet    - Examines the Go code for potential issues"
	@echo "  lint   - Runs a linter on the Go code"
