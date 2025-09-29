.PHONY: run test build test-setup

run:
	go run ./cmd/asciiart ./testdata/test.png

test-setup:
	go run ./cmd/genimage

test: test-setup
	go test ./... -v

build:
	go build -o bin/asciiart ./cmd/asciiart
