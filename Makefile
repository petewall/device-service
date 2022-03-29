SHELL = /bin/bash
GO_VERSION := $(shell go version)
GO_VERSION_REQUIRED = go1.17
GO_VERSION_MATCHED := $(shell go version | grep $(GO_VERSION_REQUIRED))
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)

# #### DEPS ####
.PHONY: deps deps-go-binary deps-golangci-lint deps-modules

deps-go-binary:
ifndef GO_VERSION
	$(error Go not installed)
endif
ifndef GO_VERSION_MATCHED
	$(error Required Go version is $(GO_VERSION_REQUIRED), but was $(GO_VERSION))
endif
	@:

deps-golangci-lint: deps-go-binary
ifndef HAS_GOLANGCI_LINT
	cd /; go get github.com/golangci/golangci-lint/cmd/golangci-lint
endif

deps-modules: deps-go-binary
	go mod download

deps: deps-modules


# #### TEST ####
.PHONY: lint

lint: deps-golangci-lint
	golangci-lint run

# #### BUILD ####
.PHONY: build
SOURCES = $(shell find . -name "*.go" | grep -v "_test\." )

build/device-service: $(SOURCES) deps
	go build -o build/device-service github.com/petewall/device-service/v2

build: build/device-service

# #### RUN ####
.PHONY: run

run: build
	./build/build/device-service
