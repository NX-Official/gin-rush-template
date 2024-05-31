
all: test

test:
	go test -v ./...

.PHONY: test