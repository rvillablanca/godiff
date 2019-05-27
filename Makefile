GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=godiff

all: build

build:
	$(GOBUILD) -v ./cmd/godiff

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN) ./...
	rm -f $(BINARY_NAME)

install:
	$(GOINSTALL) ./cmd/godiff

release:
	rm -rf ./dist
	goreleaser