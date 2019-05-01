.DEFAULT_GOAL := help
Owner := youyo
Name := gaws
Repository := "github.com/$(Owner)/$(Name)"
GithubToken := ${GITHUB_TOKEN}
Version := $(shell git describe --tags --abbrev=0)

## Setup
setup:
	GO111MODULE=off go get -v github.com/Songmu/goxz/cmd/goxz
	GO111MODULE=off go get -v github.com/tcnksm/ghr
	GO111MODULE=off go get -v github.com/jstemmer/go-junit-report

## Run tests
test:
	go test -v -cover \
		$(shell go list ./...)

## Execute `go run`
run:
	go run main.go ${OPTION}

## Vendoring
vendoring:
	go mod download

## Build
build:
	go get
	goxz -os=darwin,linux -arch=amd64 -d=pkg

## Release
release:
	ghr -t ${GithubToken} -u $(Owner) -r $(Name) --replace $(Version) pkg/

## update homebrew
update-homebrew:
	curl -s -X POST \
		-u ${CIRCLE_API_TOKEN}: \
		-d build_parameters[GAWS_VERSION]=$(Version) \
		-d build_parameters[CIRCLE_JOB]=release \
		https://circleci.com/api/v1.1/project/github/youyo/homebrew-gaws/tree/master

## Remove packages
clean:
	rm -rf pkg/

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: help
.SILENT:
