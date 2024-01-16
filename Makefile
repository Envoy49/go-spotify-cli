# Dependency versions
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
	go test -race -v ./..

##@ Dependencies

deps: bin/goreleaser 
deps: ## Install dependencies

bin/goreleaser:
	go install github.com/sigstore/cosign/v2/cmd/cosign@latest
	curl -sfL https://goreleaser.com/static/run | VERSION=v${GORELEASER_VERSION} bash -s -- --version