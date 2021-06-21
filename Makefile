GOENV = CGO_ENABLED=0
GO = $(GOENV) go

all: mod tools test lint build

mod:
	$(GO) mod download

mod-tidy:
	$(GO) mod tidy

tools:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint

test:
	$(GO) test ./...

lint:
	$(GOENV) golangci-lint run

build:
	$(GO) build -o bin/

.PHONY: all mod mod-tidy tools test lint build
