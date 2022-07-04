.PHONY: build

default: build

BINARY=go-proj-boot
GIT_VERSION := $(shell git rev-parse HEAD)
GO_VERSION := $(shell go version)
BUILD_TIME := $(shell date +%FT%T%z)

LDFLAGS=-ldflags '-s -X "github.com/fs714/go-proj-boot/global.GitVersion=${GIT_VERSION}" -X "github.com/fs714/go-proj-boot/global.GoVersion=${GO_VERSION}" -X "github.com/fs714/go-proj-boot/global.BuildTime=${BUILD_TIME}"'

build:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${BINARY} ${LDFLAGS}
clean:
	rm -rf bin/
