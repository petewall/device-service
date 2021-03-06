GO_VERSION := $(shell go version)
GO_VERSION_REQUIRED = go1.18
GO_VERSION_MATCHED := $(shell go version | grep $(GO_VERSION_REQUIRED))
HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_COUNTERFEITER := $(shell command -v counterfeiter;)

# #### DEPS ####
.PHONY: deps-go-binary deps-counterfeiter deps-golangci-lint deps-modules

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
	go install github.com/maxbrunsfeld/counterfeiter/v6@latest
endif

deps-ginkgo: deps-go-binary
ifndef HAS_GINKGO
	go install github.com/onsi/ginkgo/ginkgo@latest
endif

deps-golangci-lint: deps-go-binary
ifndef HAS_GOLANGCI_LINT
ifeq ($(PLATFORM), Darwin)
	brew install golangci-lint
endif
ifeq ($(PLATFORM), Linux)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
endif
endif

deps-modules: deps-go-binary
	go mod download

# #### SRC ####
lib/libfakes/fake_dbinterface.go: lib/db.go
	go generate lib/db.go

# #### TEST ####
.PHONY: lint

lint: deps-golangci-lint
	golangci-lint run

test: lib/libfakes/fake_dbinterface.go deps-modules deps-ginkgo
	ginkgo -r -skipPackage test .

integration-test: deps-modules deps-ginkgo
	ginkgo -r test/integration

test-all: lib/libfakes/fake_dbinterface.go deps-modules deps-ginkgo
	ginkgo -r .

# #### BUILD ####
.PHONY: build
SOURCES = $(shell find . -name "*.go" | grep -v "_test\." )

build/device-service: $(SOURCES) deps-modules
	go build -o build/device-service github.com/petewall/device-service/v2

build: build/device-service

build-image:
	docker build -t petewall/device-service .

# #### RUN ####
.PHONY: run

run: build/device-service
	./build/device-service
