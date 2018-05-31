all: build

build:
	go build -o changelog main.go

test:
	go test ./...
