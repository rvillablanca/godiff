GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=godiff

all: build

build:
	$(GOBUILD) -v .

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

platforms:
	rm -rf dist
	mkdir dist
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o dist/godiff.exe ./cmd/godiff
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o dist/godiff-linux ./cmd/godiff
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o dist/godiff-darwin ./cmd/godiff
