
.PHONY: build fmt test run bench

default: build

build: fmt
	go build

run: build
	./april -consume 3 -produce 5

test:
	go test

bench:
	 go test -v -test.bench=".*" -count=5

fmt:
	go fmt ./...
