generate:
	@go generate ./...
.PHONY: generate

build: generate
	@echo "====> Build gcommitter cli"
	@go build -o ./bin/gcommitter main.go
.PHONY: build

release:
	@echo "====> Build and release"
	@go get github.com/goreleaser/goreleaser
	@goreleaser
.PHONY: release

test:
	@go test ./...
.PHONY: test

test.cov:
	@go test ./... -coverprofile=coverage.txt -covermode=atomic
.PHONY: test.cov
