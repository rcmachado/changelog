all: build

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build: deps
	go build -o changelog main.go

test:
	go test -cover ./...

lint:
	[[ `which gometalinter.v2` ]] || go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
	gometalinter.v2 --vendor ./...

release: build
	./changelog release $(V) -o CHANGELOG.md
