
all: test

test:
	go test -v ./...

lint:
	golangci-lint run

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.0

.PHONY: test