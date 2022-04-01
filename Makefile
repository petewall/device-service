SHELL = /bin/bash
GO_VERSION := $(shell go version)
GO_VERSION_REQUIRED = go1.17
GO_VERSION_MATCHED := $(shell go version | grep $(GO_VERSION_REQUIRED))
HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_COUNTERFEITER := $(shell command -v counterfeiter;)

# #### DEPS ####
.PHONY: deps deps-go-binary deps-counterfeiter deps-golangci-lint deps-modules

deps-go-binary:
ifndef GO_VERSION
	$(error Go not installed)
endif
ifndef GO_VERSION_MATCHED
	$(error Required Go version is $(GO_VERSION_REQUIRED), but was $(GO_VERSION))
endif
	@:

deps-counterfeiter: deps-go-binary
ifndef HAS_COUNTERFEITER
	cd /; go get github.com/maxbrunsfeld/counterfeiter/v6
endif

deps-ginkgo: deps-go-binary
ifndef HAS_GINKGO
	cd /; go get github.com/onsi/ginkgo/ginkgo github.com/onsi/gomega
endif

deps-golangci-lint: deps-go-binary
ifndef HAS_GOLANGCI_LINT
	cd /; go get github.com/golangci/golangci-lint/cmd/golangci-lint
endif

deps-modules: deps-go-binary
	go mod download

deps: deps-modules deps-counterfeiter deps-ginkgo


# #### TEST ####
.PHONY: lint

lint: deps-golangci-lint
	golangci-lint run

test: deps-counterfeiter deps-ginkgo
	ginkgo -r .

# #### BUILD ####
.PHONY: build
SOURCES = $(shell find . -name "*.go" | grep -v "_test\." )

build/device-service: $(SOURCES) deps
	go build -o build/device-service github.com/petewall/device-service/v2

build: build/device-service

# #### RUN ####
.PHONY: run

run: build
	./build/device-service
