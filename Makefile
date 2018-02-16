.PHONY: fmt
fmt:
	gofmt -d -l -w ./ice ./playground

.PHONY: test-playground
test-playground:
	go test -v ./playground/...