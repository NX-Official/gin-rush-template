all: test lint

test:
	go test -v ./...

lint:
	npx commitlint --from HEAD~1 --to HEAD --verbose
	golangci-lint run -v --timeout 5m


lint-install:
	npm install conventional-changelog-conventionalcommits
	npm install commitlint@latest
	npm i -D @xyclos/commitlint-plugin-references
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.0

.PHONY: test lint lint-install