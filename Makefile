GIT_REPO = github.com/at15/go.ice
VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)

.PHONY: sync-local
sync-local:
	python script/sync_local.py

.PHONY: fmt
fmt:
	gofmt -d -l -w ./ice ./playground ./_example

.PHONY: test-playground
test-playground:
	go test -v ./playground/...

#--- example ----
.PHONY: install-icehub
install-icehub:
	cd example/github; make install
#--- end of example ----