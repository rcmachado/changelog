all: build

build:
	go build -o changelog main.go

test:
	go test -cover ./...

lint: LINT_BIN=$(shell go env GOPATH)/bin
lint:
	[ `which $(LINT_BIN)/golangci-lint` ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LINT_BIN) v1.22.2
	$(LINT_BIN)/golangci-lint run

release: build
	./changelog release $(V) -o CHANGELOG.md
