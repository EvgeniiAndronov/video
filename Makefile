.PHONY: build
build:
	go build -v ./cmd/videoserver

.PHONY: test
test:
	go test -v -race ./...

.DEFAULT_GOAL := build