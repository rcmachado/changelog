GOBIN?=$(shell go env GOPATH)/bin

.PHONY: install
install: ## Install tools used by the project
	fgrep '_' tools.go | cut -f2 -d' ' | xargs go install
	# golangci-lint project doesn't recommend to install from go modules
	[ `which $(GOBIN)/golangci-lint` ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LINT_BIN) v1.22.2

build:
	go build -o changelog main.go

test:
	go test -cover ./...

lint:
	$(GOBIN)/golangci-lint run

release: build
	./changelog release $(V) -o CHANGELOG.md

.DEFAULT_GOAL:=help
.PHONY: help
help: COLUMN_SIZE=15
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-$(COLUMN_SIZE)s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
