.PHONY: fmt
fmt:
	gofmt -d -l -w ./ice ./playground

# --- test ---
.PHONY: test
test:
	go test -v -cover ./ice/...

.PHONY: test-cover
test-cover:
	go test -coverprofile=coverage.txt -covermode=atomic ./ice/...

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