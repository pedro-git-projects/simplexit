BINARY  := simplexit
MODULE  := nilptr.dev/simplexit
DIST    := dist

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w

.PHONY: all build release run fmt vet clean

all: build

build:
	go build -o $(BINARY) .

release:
	mkdir -p $(DIST)
	CGO_ENABLED=1 go build -trimpath -ldflags "$(LDFLAGS)" -o $(DIST)/$(BINARY) .

run: build
	./$(BINARY)

fmt:
	gofmt -w .

vet:
	go vet ./...

clean:
	rm -f $(BINARY)
	rm -rf $(DIST)
