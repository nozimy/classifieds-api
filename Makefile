build:
	go build -v -o ./classifieds-api ./cmd/classifieds-api

run:
	go run ./cmd/classifieds-api/main.go

test:
	go test -v -cover -race -timeout 30s ./...

.DEFAULT_GOAL := build

.PHONY: build run test