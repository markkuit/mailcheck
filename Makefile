LASTTAG := $(shell git describe --abbrev=0 --tags 2>/dev/null)

GOCMD=go
GOBUILD=$(GOCMD) build
LDFLAGS=-X github.com/markkuit/mailcheck/internal/commons.Version=$(LASTTAG)
BINDIR=$(CURDIR)/bin
BINNAME=mailcheck

default: build run
build:
	$(GOBUILD) -ldflags "$(LDFLAGS)" -v -o $(BINDIR)/$(BINNAME) cmd/mailcheck/mailcheck.go
run:
	$(BINDIR)/$(BINNAME) $(ARGS)
