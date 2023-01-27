.PHONY: build clean test scan

install-dev:
	go install honnef.co/go/tools/cmd/staticcheck@latest

test: ## Run unittests
	go test ./...
test-out:
	mkdir -p report
	go test -json ./... | tee report/test.out
coverage: ## Describes how much of a package's code is exercised by running the package's tests
	go test ./... -json -cover
coverage-out:
	mkdir -p report
	go test ./... -json -cover -coverprofile=report/coverage.out
coverage-html-out:
	go tool cover -html=report/coverage.out
vet: ## Reports suspicious constructs
	go vet ./...
vet-out:
	mkdir -p report
	go vet ./... 2>&- | tee report/govet.out

staticcheck-out:
	mkdir -p report
	staticcheck -checks=all ./... | tee report/staticcheck.out

scan: test test-out coverage-out vet-out staticcheck-out
scan-cli: test coverage vet staticcheck

see-coverage:
	go tool cover -html report/coverage.out -o cover.html
	open cover.html

go-generate:
	go generate ./...

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
