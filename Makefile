GO ?= go

build:
	@echo "====> Build gen"
	@$(GO) build -o ./bin/gct *.go
.PHONY: build

install.dev:
	@$(GO) get -u github.com/golang/dep/cmd/dep
	@$(GO) get -u github.com/jteeuwen/go-bindata/...
	@dep ensure
.PHONY: install.dev

bindata:
	@go-bindata -o bindata/bindata.go -pkg bindata template
.PHONY: bindata

release:
	@echo "====> Build and release"
	@curl -sL https://git.io/goreleaser | bash
.PHONY: release

