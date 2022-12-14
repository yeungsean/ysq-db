
OS := $(shell uname -s)
GOPATH ?= $(shell go GOPATH)
GOROOT ?= $(shell go GOROOT)
GOBIN ?= $(GOROOT)/bin
GO ?= go
GOLANGCI-LINT ?= golangci-lint
DELVE ?= dlv
MAIN = main.go

BUILD_DATE ?= $(shell date -u)
BUILD_HASH ?= $(shell git rev-parse HEAD)

DIR = $(shell dirname `pwd`)
APP = $(shell basename `pwd`)
PKG = github.com/yeungsean/ysq-db/pkg

vet:
	@$(GO) vet ./...

test:
	@$(GO) test -gcflags=all=-l ./...

mod:
	@$(GO) mod tidy
	git diff --exit-code go.mod
