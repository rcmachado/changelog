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
	$(LINT_BIN)/golangci-lint run

release: build
	./changelog release $(V) -o CHANGELOG.md
