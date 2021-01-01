NAME = steve
GOENV = CGO_ENABLED=0
GO = $(GOENV) go
KIND = KIND_CLUSTER_NAME=$(NAME) kind

all: mod tools test-all build
.PHONY: all

mod:
	$(GO) mod download
mod-tidy:
	$(GO) mod tidy
.PHONY: mod mod-tidy

tools: 
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint
.PHONY: tools 

test:
	$(GO) test -coverprofile=coverage.out -v ./...
lint:
	$(GOENV) golangci-lint run
test-all: test lint
.PHONY: test lint test-all

build:
	$(GO) build -o ./bin/$(NAME) ./cmd/$(NAME)
.PHONY: build

clean:
	$(GO) clean
	golangci-lint cache clean
	rm -rf bin coverage.out
.PHONY: clean
