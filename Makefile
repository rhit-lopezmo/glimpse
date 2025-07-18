.DEFAULT_GOAL := build

.PHONY: fmt vet build run clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build .

run: vet
	go run .

clean:
	go clean
