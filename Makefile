# --- packages ---
PKGST=./api ./cli ./cmd ./db ./httpclient ./udash
PKGS=./httpclient/...
# --- packages ---

.PHONY: install
install: fmt
	go install ./cmd/dk

.PHONY: fmt
fmt:
	goimports -d -l -w $(PKGST)

# --- test ---
.PHONY: test test-cover test-cover-html

test:
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
.PHONY: dep-install
	dep ensure -v
.PHONY: dep-update
dep-update:
	dep ensure -update -v
# --- dependency management ---

.PHONY: loc
loc:
	tokei .