## Basic
PROJ="magikarp"
TIME=$$(date +%Y-%m-%d_%H:%M)

## Git
COMMIT=$$(git rev-parse --short HEAD)
VERSION=$$( if $$(git describe --abbrev=0 > /dev/null 2>&1); then echo $$(git describe --abbrev=0) ; else echo $(COMMIT) ; fi )

## CMD
GO := go

##@ Basic

version:  ## Show version
	echo ${VERSION}

.PHONY: version


##@ Lint

install-lint:  ## Install golangci-lint binary to ./bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.31.0

lint-testing:  ## Run lint
	./bin/golangci-lint run ./...

.PHONY: install-lint lint-testing

##@ Go CMD

run:  ## Run
	${GO} run cmd/*.go

run-dev:  ## Run as developer mode
	${GO} run cmd/*.go --log_level debug

build:  ## Go build bin
	${GO}  build -o ./bin/${PROJ}  -ldflags "-s -w -X main.date=${TIME} -X main.commit=${COMMIT} -X main.version=${VERSION}" ./cmd/*.go

##@ Testing

test:  ## Run go test
	${GO} test -v -cover ./...


.PHONY: run run-dev build

##@ Dev testing

test-service-up:  ## Service up for dev testing.
	docker-compose -f ./docker-compose.test.yaml up -d

install-pgcli:  ## Install pgcli with pip
	pip install pgcli

pgcli-conn:  ## Conn to postgres with pgcli
	pgcli -h localhost -p 5432 -U postgres

.PHONY: test-service-up install-pgcli pgcli-conn


##@ Help

.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
