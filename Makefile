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
ICEHUB_PKG = github.com/at15/go.ice/_example/github/pkg
ICEHUB_VERSION = 0.0.1
ICEHUB_FLAGS = -X $(ICEHUB_PKG)/common.version=$(ICEHUB_VERSION) -X $(ICEHUB_PKG)/common.gitCommit=$(BUILD_COMMIT) -X $(ICEHUB_PKG)/common.buildTime=$(BUILD_TIME) -X $(ICEHUB_PKG)/common.buildUser=$(CURRENT_USER)
.PHONY: install-icehub
install-icehub:
	go install -ldflags "$(ICEHUB_FLAGS)" ./_example/github/cmd/icehubctl
	go install -ldflags "$(ICEHUB_FLAGS)" ./_example/github/cmd/icehubd
#--- end of example ----