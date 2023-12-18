.PHONY: default run build test clean

APP_NAME=mangad

default: run

run:
	@go run ./cmd/mangad/main.go

build:
	@go build -o $(APP_NAME) ./cmd/mangad/main.go

test:
	@go test ./...

clean:
	@rm -f $(APP_NAME)