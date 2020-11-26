
VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT}"
MODFLAGS=-mod=vendor
TESTFLAGS=-cover

PLATFORMS:=darwin linux windows

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

.PHONY: all dev dist $(PLATFORMS) clean test
