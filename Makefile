.DEFAULT_GOAL := build

.PHONY: fmt vet build clean tidy path

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

tidy: 
	go mod tidy

build: tidy vet
	go build .

clean:
	go clean

path:
	sudo cp glimpse /usr/local/bin/

