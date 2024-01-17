# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

export PATH := $(abspath bin/):${PATH}

# Dependency versions
GOLANGCI_VERSION = 1.53.3
GORELEASER_VERSION = 1.18.2

##@ General

# Targets commented with ## will be visible in "make help" info.
# Comments marked with ##@ will be used as categories for a group of targets.

.PHONY: help
.DEFAULT_GOAL := help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.PHONY: build
build: ## Build binary
	@mkdir -p build
	go build -race -o build/go-spotify-cli .

.PHONY: artifacts
artifacts: binary-snapshot
artifacts: ## Build artifacts

.PHONY: binary-snapshot
binary-snapshot: ## Build binary snapshot
	goreleaser --snapshot --skip=publish --clean --snapshot

##@ Checks

.PHONY: test
test: ## Run tests
	go test -race -v ./

.PHONY: lint
lint: lint-go
lint: ## Run linters

.PHONY: lint-go
lint-go:
	golangci-lint run $(if ${CI},--out-format github-actions,)

.PHONY: fmt
fmt: ## Format code
	golangci-lint run --fix

##@ Dependencies

deps: golangci-lint goreleaser
deps: ## Install dependencies

golangci-lint:
	@mkdir -p bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- v${GOLANGCI_VERSION}

goreleaser:
	@mkdir -p bin
	go install github.com/sigstore/cosign/v2/cmd/cosign@latest
	curl -sfL https://goreleaser.com/static/run | VERSION=v${GORELEASER_VERSION} bash -s -- --version