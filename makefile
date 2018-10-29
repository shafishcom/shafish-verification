# tokuhash
export GOPATH := $(shell pwd)
export PATH := $(GOPATH)/bin:$(PATH)

build:
	@echo "--> Building..."
	go build -v -o bin/shafish-verification  src/main.go

clean:
	@echo "--> Cleaning..."
	@go clean

test:
	go get -v github.com/stretchr/testify/assert
	@echo "--> Testing..."
	@$(MAKE) testverify

testverify:
	go test -v -race verify

pkgs =	verify

fmt:
	go vet $(pkgs)
	gofmt -s -w ./

coverage:
	go get -v github.com/pierrre/gotestcover
	gotestcover -coverprofile=coverage.out -v $(pkgs)
	go tool cover -html=coverage.out

.PHONY: clean test coverage
