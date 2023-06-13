SHELL := /bin/bash
GREEN := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
RESET := $(shell tput -Txterm sgr0)

.PHONY: help

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "${YELLOW}%-16s${GREEN}%s${RESET}\n", $$1, $$2}' $(MAKEFILE_LIST)

tests: ## Run the tests of the project.
	go test ./... -v

coverage: ## Generate code coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage-report.html

deps: ## Download dependencies
	go mod tidy
	go mod download

build: ## Build binary for local operating system
	go generate ./...
	go build .

clean: ## Remove build related file
	go clean
