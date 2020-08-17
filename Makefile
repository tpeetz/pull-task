# Go parameters
GOPATH=$(HOME)/go
GOBIN=$(HOME)/go/bin
GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOCOVER=$(GOCMD) tool cover
COBERTURA=$(GOBIN)/gocover-cobertura
GOLINT=$(GOBIN)/golint
GOVET=$(GOCMD) vet

# Project parameters
GONAME=pull-task
GOFILE=cmd/pull-task/main.go

.PHONY: all deps clean clean-bin clean-doc

all: deps build  build-win

deps:
	GOPATH=$(GOPATH) $(GOGET) -u github.com/spf13/cobra/cobra
	GOPATH=$(GOPATH) $(GOGET) -u github.com/jstemmer/go-junit-report
	GOPATH=$(GOPATH) $(GOGET) -u github.com/inconshreveable/mousetrap
	GOPATH=$(GOPATH) $(GOGET) -u github.com/mitchellh/go-homedir
	GOPATH=$(GOPATH) $(GOGET) -u github.com/t-yuki/gocover-cobertura
	GOPATH=$(GOPATH) $(GOGET) -u golang.org/x/lint/golint

build: bin/$(GONAME)

build-win: bin/$(GONAME).exe

bin/$(GONAME): $(GOFILE)
	@echo "Building $(GOFILE) to ./bin"
	@GOPATH=$(GOPATH) $(GOBUILD) -v -ldflags="-X main.version=$(shell git describe --tags)" -o bin/$(GONAME) $(GOFILE)

bin/$(GONAME).exe: $(GOFILE)
	@echo "Building $(GOFILE) to ./bin"
	env GOOS=windows GOARCH=amd64 GOPATH=$(GOPATH) $(GOBUILD) -v -ldflags="-X main.version=$(shell git describe --tags)" -o bin/$(GONAME).exe $(GOFILE)

install:
	@echo using $(GOPATH)
	GOPATH=$(GOPATH) $(GOCMD) install -v -ldflags="-X main.version=$(shell git describe --tags)" ./...

check: test coverage lint vet

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=cover.out ./...
	$(GOCOVER) -html=cover.out -o coverage.html
	$(COBERTURA) < cover.out > coverage.xml

lint:
	$(GOLINT) ./...

vet:
	$(GOVET) -v ./...

doc:
	cd docs; ./gradlew build

clean: clean-doc clean-bin


clean-doc:
	@echo "Cleaning Gradle build"
	cd docs; ./gradlew clean

clean-bin:
	@echo "Cleaning Go build"
	rm -rf bin/

# Cross compilation
#build-linux:
#       CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
#docker-build:
#       docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
#
