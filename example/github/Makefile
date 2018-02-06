PKG = github.com/at15/go.ice/example/github/pkg
VERSION = 0.0.1
FLAGS = -X $(PKG)/common.version=$(VERSION) -X $(PKG)/common.gitCommit=$(BUILD_COMMIT) -X $(PKG)/common.buildTime=$(BUILD_TIME) -X $(PKG)/common.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/icehubctl
	go install -ldflags "$(FLAGS)" ./cmd/icehubd

.PHONY: fmt
fmt:
	gofmt -d -l -w ./pkg ./cmd

.PHONY: keygen
keygen:
	openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout icehub.key -out icehub.crt -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=someone@random.com"
