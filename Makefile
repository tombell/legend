NAME=legend
ENTRY=./cmd/$(NAME)
OUTPUT=dist/$(NAME)

VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT)"
MODFLAGS=-mod=vendor
TESTFLAGS=-cover

CGO_LDFLAGS:=-L$(shell brew --prefix openssl)/lib
CGO_CPPFLAGS:=-I$(shell brew --prefix openssl)/include

PLATFORMS:=darwin

all: dev

dev: build-web
	@echo building $(OUTPUT)...
	@CGO_LDFLAGS=$(CGO_LDFLAGS) CGO_CPPFLAGS=$(CGO_CPPFLAGS) \
		go build $(MODFLAGS) $(LDFLAGS) -o $(OUTPUT) $(ENTRY)

dist: build-web $(PLATFORMS)

$(PLATFORMS):
	@echo building $(OUTPUT)-$@-amd64...
	@CGO_LDFLAGS=$(CGO_LDFLAGS) CGO_CPPFLAGS=$(CGO_CPPFLAGS) GOOS=$@ GOARCH=amd64 \
		go build $(MODFLAGS) $(LDFLAGS) -o $(OUTPUT)-$@-amd64 $(ENTRY)

run:
	@$(OUTPUT)

watch:
	@while sleep 1; do \
		trap "exit" SIGINT; \
		find . \
			-type d \( -name vendor -o -name node_modules \) -prune -false -o \
			-type f \( -name "*.go" -o -name "*.tsx" -o -name "*.ts" \) \
			| entr -c -d -r make dev run; \
	done

build-web:
	@npm run build

lint-web:
	@npm run lint

test:
	@go test ${MODFLAGS} ${TESTFLAGS} ./...

clean:
	@rm -fr dist pkg/web/public/app.js

modules:
	@CGO_LDFLAGS=$(CGO_LDFLAGS) CGO_CPPFLAGS=$(CGO_CPPFLAGS) \
		go get -u ./... && go mod download && go mod tidy && go mod vendor

.PHONY: all dev dist $(PLATFORMS) run watch build-web lint-web test clean modules
