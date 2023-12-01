LINTER := $(shell golangci-lint --version)

.PHONY: lint

lint:
    ifeq ($(LINTER),)
        GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.2
    endif
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix

build:
	docker build -t beanflow .
