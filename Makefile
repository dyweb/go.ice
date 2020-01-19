GOMOD = GO111MODULE=on go mod

# --- build vars ---
VERSION = 0.0.2
BUILD_COMMIT := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S%z)
DOCKER_REPO = dyweb/go.ice
# --- build vars ---

# --- packages ---
PKGST=./cli
PKGS=./cli/...
# --- packages ---

.PHONY: install
install: fmt test
	go install ./cmd/dk

.PHONY: fmt
fmt:
	goimports -d -l -w $(PKGST)

# --- test ---
.PHONY: test test-verbose test-cover test-cover-html

test:
	go test -cover $(PKGS)

test-verbose:
	go test -v -cover $(PKGS)

test-cover:
# https://github.com/codecov/example-go
	go test -coverprofile=coverage.txt -covermode=atomic $(PKGS)

test-cover-html: test-cover
	go tool cover -html=coverage.txt

.PHONY: test-playground
test-playground:
	go test -v ./playground/...
# --- test ---

.PHONY: generate
generate:
	gommon generate -v

# --- dependency management ---
.PHONY: dep-install dep-update mod-init mod-update
dep-install:
	dep ensure -v
dep-update:
	dep ensure -update -v
mod-init:
	$(GOMOD) init
mod-update:
	$(GOMOD) tidy
# --- dependency management ---

# --- docker ---
.PHONY: docker-build docker-push

docker-build:
	docker build -t $(DOCKER_REPO):$(VERSION) .

docker-push:
	docker push $(DOCKER_REPO):$(VERSION)
# --- docker ---

.PHONY: loc
loc:
	tokei .