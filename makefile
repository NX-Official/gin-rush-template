all: test lint

test:
	go test -v ./...

lint:
	golangci-lint run --timeout 5m

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.0

.PHONY: test lint lint-install