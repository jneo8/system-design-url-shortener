GO := go

##@ Lint

install-lint:  ## Install golangci-lint binary to ./bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.31.0

lint-testing:  ## Run lint
	./bin/golangci-lint run ./...

.PHONY: install-lint lint-testing

##@ Go CMD

run:  ## Run
	${GO} run cmd/*.go

.PHONY: run

##@ Help

.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
