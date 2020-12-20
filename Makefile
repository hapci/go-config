LINTER_VERSION := v1.33.0

check: test lint

setup: check-env ## install required development utilities
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin $(LINTER_VERSION)

lint: check-env ## run lint check
	$(GOPATH)/bin/golangci-lint run --fix

test: ## run the unit tests
	go test ./... -race -coverprofile .test_coverage.txt

test-coverage: test ## show test coverage
	go tool cover -html=.test_coverage.txt

check-env:
ifeq ($(GOPATH),)
	@echo "GOPATH is missing (set GOPATH)"
	@exit 1
endif

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort
