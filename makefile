# Advent of code day
DAY=1

# Things you don't want to change
REPO_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
GOLINT_PATH := $(REPO_DIR)/bin/golangci-lint

.PHONY: help run lint test
.DEFAULT_GOAL := help

help: ## 💬 This help message :)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

lint: ## 🌟 Lint & format
	@$(GOLINT_PATH) > /dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh
	$(GOLINT_PATH) run ./...

run: ## 🚀 Run code for given day
	cd day-$(DAY); go run main.go

test: ## 🧪 Test the code for a given day
	cd day-$(DAY); go test -v