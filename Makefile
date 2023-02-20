SHELL := /bin/bash

HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_COUNTERFEITER := $(shell command -v counterfeiter;)
PLATFORM := $(shell uname -s)

# #### DEPS ####
.PHONY: deps-counterfeiter deps-ginkgo deps-modules

deps-counterfeiter:
ifndef HAS_COUNTERFEITER
	go install github.com/maxbrunsfeld/counterfeiter/v6@latest
endif

deps-ginkgo:
ifndef HAS_GINKGO
	go install github.com/onsi/ginkgo/v2/ginkgo
endif

deps-modules:
	go mod download

# #### SRC ####
internal/internalfakes/fake_dbinterface.go: internal/db.go
	go generate internal/db.go

# #### TEST ####
.PHONY: lint test-units integration-test test

lint: internal/internalfakes/fake_dbinterface.go deps-modules
ifndef HAS_GOLANGCI_LINT
ifeq ($(PLATFORM), Darwin)
	brew install golangci-lint
endif
ifeq ($(PLATFORM), Linux)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
endif
endif
	golangci-lint run

test-units: internal/internalfakes/fake_dbinterface.go deps-modules deps-ginkgo
	ginkgo -r -skip-package test .

test-features: deps-modules deps-ginkgo
	ginkgo -r test

test: lint test-units test-features

# #### BUILD ####
.PHONY: build
SOURCES = $(shell find . -name "*.go" | grep -v "_test\." )

build/device-service: $(SOURCES) deps-modules
	go build -o build/device-service github.com/petewall/device-service

build: build/device-service

build-image:
	docker build -t petewall/device-service .

# #### RUN ####
.PHONY: run

run: build/device-service
	./build/device-service
