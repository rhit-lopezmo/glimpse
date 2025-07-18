.DEFAULT_GOAL := build

.PHONY: fmt vet build clean init tidy path

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

init:
	@if [ ! -f go.mod ]; then \
		echo "Initializing go.mod..."; \
		go mod init glimpse; \
	else \
		echo "go.mod already exists, skipping init."; \
	fi

tidy: 
	go mod tidy

build: init tidy vet
	go build .

clean:
	go clean

path:
	sudo cp glimpse /usr/local/bin/

