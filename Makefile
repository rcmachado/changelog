all: build

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build: deps
	go build -o changelog main.go

test:
	go test -cover ./...

release: build
	./changelog release $(V) -o CHANGELOG.md
