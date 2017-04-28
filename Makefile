build:
	go build -i -ldflags "-w -s -extldflags=-static"

test:
	go test -i -race
	go test -cover -race $(shell go list ./... | grep -v vendor/)

.PHONY: test
