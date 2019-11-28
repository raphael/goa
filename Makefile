#! /usr/bin/make
#
# Makefile for Goa v3
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
# - "release" creates a new release commit, tags the commit and pushes the tag to GitHub.
#   "release" also updates the examples and plugins repo and pushes the updates to GitHub.
#
# Meta targets:
# - "all" is the default target, it runs "lint" and "test"
#
MAJOR = 3
MINOR = 0
BUILD = 8

GOOS=$(shell go env GOOS)
GO_FILES=$(shell find . -type f -name '*.go')

ifeq ($(GOOS),windows)
EXAMPLES_DIR="$(GOPATH)\src\goa.design\examples"
PLUGINS_DIR="$(GOPATH)\src\goa.design\plugins"
GOBIN="$(GOPATH)\bin"
else
EXAMPLES_DIR=$(GOPATH)/src/goa.design/examples
PLUGINS_DIR=$(GOPATH)/src/goa.design/plugins
GOBIN=$(GOPATH)/bin
endif

# Only list test and build dependencies
# Standard dependencies are installed via go get
DEPEND=\
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/golang/protobuf/proto \
	honnef.co/go/tools/cmd/staticcheck

all: lint test

travis: depend all #test-examples test-plugins

# Install protoc
PROTOC_VERSION=3.6.1
ifeq ($(GOOS),linux)
PROTOC=protoc-$(PROTOC_VERSION)-linux-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
else
	ifeq ($(GOOS),darwin)
PROTOC=protoc-$(PROTOC_VERSION)-osx-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
	else
		ifeq ($(GOOS),windows)
PROTOC=protoc-$(PROTOC_VERSION)-win32
PROTOC_EXEC="$(PROTOC)\bin\protoc.exe"
		endif
	endif
endif
depend:
	@go get -v $(DEPEND)
	@env GO111MODULE=off go get github.com/hashicorp/go-getter/cmd/go-getter && \
		go-getter https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC).zip $(PROTOC) && \
		cp $(PROTOC_EXEC) $(GOBIN) && \
		rm -r $(PROTOC) && \
		echo "`protoc --version`"
	@go get -t -v ./...

lint:
	@if [ "`goimports -l $(GO_FILES) | tee /dev/stderr`" ]; then \
		echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
	fi
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi
	@if [ "`staticcheck -checks all ./... | grep -v ".pb.go" | tee /dev/stderr`" ]; then \
		echo "^ - staticcheck errors!" && echo && exit 1; \
	fi

test:
	env GO111MODULE=on go test ./...

release:
	@git diff-index --quiet HEAD
	go mod tidy
	sed -i '' 's/Build = \[0-9]+\/Build = $(BUILD)' pkg/version.go
	sed -i '' 's/Current Release: `v3.\.*\/Current Release: `v$(MAJOR).$(MINOR).$(BUILD)`' README.md
	git add .
	git commit -m "Release v$(MAJOR).$(MINOR).$(BUILD)"
	git push origin v$(MAJOR)
	git tag v$(MAJOR).$(MINOR).$(BUILD)
	git push origin v$(MAJOR).$(MINOR).$(BUILD)
	cd $(GOPATH)/src/goa.design/examples && \
		git checkout master && \
		git diff-index --quiet HEAD && \
		sed -i '' 's/goa.design\/goa\/v$(MAJOR) v$(MAJOR)\.*\/goa.design\/goa\/v$(MAJOR) v$(MAJOR).$(MINOR).$(BUILD)' go.mod && \
		make && \
		git add . && \
		git commit -m "Release v$(MAJOR).$(MINOR).$(BUILD)" && \
		git push origin master && \
		git tag v$(MAJOR).$(MINOR).$(BUILD) && \
		git push origin v$(MAJOR).$(MINOR).$(BUILD)
	cd $(GOPATH)/src/goa.design/plugins && \
		git checkout v$(MAJOR) && \
		git diff-index --quiet HEAD && \
		sed -i '' 's/goa.design\/goa\/v$(MAJOR) v$(MAJOR)\.*\/goa.design\/goa\/v$(MAJOR) v$(MAJOR).$(MINOR).$(BUILD)' go.mod && \
		make && \
		git add . && \
		git commit -m "Release v$(MAJOR).$(MINOR).$(BUILD)" &&\
		git push origin v$(MAJOR) && \
		git tag v$(MAJOR).$(MINOR).$(BUILD) && \
		git push origin v$(MAJOR).$(MINOR).$(BUILD)
	echo DONE RELEASING v$(MAJOR).$(MINOR).$(BUILD)!


