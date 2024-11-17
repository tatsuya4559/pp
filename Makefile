.DEFAULT_GOAL := help

pp: *.go
	go build ./...

build: pp ## Build the pp command

.PHONY: install
install: ## Install the pp command
	go install

.PHONY: test
test: ## Run tests
	go test -v ./...

.PHONY: help
help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
