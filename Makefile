.PHONY: test
test:
	go test -v -cover ./ice/...

.PHONY: test-cover
test-cover:
	go test -coverprofile=coverage.txt -covermode=atomic ./ice/...

.PHONY: loc
loc:
	cloc --exclude-dir=vendor,.idea,playground,vagrant,node_modules,example .

.PHONY: fmt
fmt:
	gofmt -d -l -w ./ice ./playground




.PHONY: test-playground
test-playground:
	go test -v ./playground/...

.PHONY: update-dep
update-dep:
	dep ensure -update