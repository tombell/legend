VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

CGO_LDFLAGS:=-L/usr/local/opt/openssl/lib
CGO_CPPFLAGS:=-I/usr/local/opt/openssl/include

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Commit=${COMMIT}"
MODFLAGS=-mod=vendor
TESTFLAGS=-cover

PLATFORMS:=darwin

all: dev

dev:
	@echo building dist/legend...
	@CGO_LDFLAGS=${CGO_LDFLAGS} CGO_CPPFLAGS=${CGO_CPPFLAGS} go build ${MODFLAGS} ${LDFLAGS} -o dist/legend ./cmd/legend

dist: $(PLATFORMS)

$(PLATFORMS):
	@echo building dist/legend-$@-amd64...
	@CGO_LDFLAGS=${CGO_LDFLAGS} CGO_CPPFLAGS=${CGO_CPPFLAGS} GOOS=$@ GOARCH=amd64 go build ${MODFLAGS} ${LDFLAGS} -o dist/legend-$@-amd64 ./cmd/legend

clean:
	@rm -fr dist/

test:
	@go test ${MODFLAGS} ${TESTFLAGS} ./...

modules:
	@CGO_LDFLAGS=${CGO_LDFLAGS} CGO_CPPFLAGS=${CGO_CPPFLAGS} go get -u ./... && go mod download && go mod tidy && go mod vendor

.PHONY: all dev dist $(PLATFORMS) clean test modules
