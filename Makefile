RELEASE = wtf
VERSION = $(shell date +%Y%m%d)
MODULE_NAME = github.com/MR5356/wtf

.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

build: clean deps ## Build the project
	go build -ldflags "-s -w -X '$(MODULE_NAME)/pkg/version.Version=$(VERSION)'" -o _output/$(RELEASE) cmd/wtf.go

deps: ## Install dependencies using go get
	go get -d -v -t ./...

clean: ## Clean
	rm -rf _output