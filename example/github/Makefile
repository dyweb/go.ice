PKG = github.com/dyweb/go.ice/example/github/pkg
VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/icehubctl
	go install -ldflags "$(FLAGS)" ./cmd/icehubd

.PHONY: clean
clean:
	rm $(shell which icehubctl)
	rm $(shell which icehubd)

.PHONY: generate
generate:
	gommon generate -v

.PHONY: fmt
fmt:
	gofmt -d -l -w ./pkg ./cmd

.PHONY: keygen
keygen:
	openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout icehub.key -out icehub.crt -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=someone@random.com"

.PHONY: dep-update
dep-update:
	dep ensure -update