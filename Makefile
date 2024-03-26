RELEASE = wtf
VERSION ?= "PROD1.1.0_R$(shell git rev-parse --short HEAD)_T$(shell date +%Y%m%d)"
MODULE_NAME = github.com/MR5356/wtf

.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

build: clean deps ## Build the project
	go build -ldflags "-s -w -X '$(MODULE_NAME)/pkg/version.Version=$(VERSION)'" -o _output/$(RELEASE) cmd/wtf.go

release: clean deps test ## Generate releases for unix systems
	chmod +x hack/release.sh
	bash -c "hack/release.sh $(VERSION) $(RELEASE) $(MODULE_NAME)"

deps: ## Install dependencies using go get
	go get -d -v -t ./...

clean: ## Clean
	rm -rf _output

test:
	go test -v ./...
