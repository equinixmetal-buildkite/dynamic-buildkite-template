GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOLINT=golangci-lint
BINARY_NAME=dynamic-buildkite-template

.PHONY: test

all: clean lint build

## lint: Runs the linter on the source
lint:
	$(GOLINT) run --fix
## build: Gathers the dependencies and builds the binary
build:
	$(GOMOD) tidy
	$(GOBUILD) -o $(BINARY_NAME) -v

## test: Runs all available tests in the source code
test:
	$(GOTEST) -v ./...

## clean: Cleans the build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

## run: Builds and runs the project
run: build
	./$(BINARY_NAME)

## rpm: Builds rpm package
rpm: build
	nfpm pkg --packager rpm --target . 
## deb: Buids deb package
deb: build
	nfpm pkg --packager deb --target .
## help: gives help instructions
help: Makefile
	@echo
	@echo "Available Commands:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo