GIT_REPO = github.com/at15/go.ice
VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)

.PHONY: sync-local
sync-local:
	python script/sync_local.py

.PHONY: fmt
fmt:
	gofmt -d -l -w ./playground
.PHONY: test-playground
test-playground:
	go test -v ./playground/...