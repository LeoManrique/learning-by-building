BINARY := $(abspath bin/cli-task-tracker)

.PHONY: build test e2e clean

build:
	go build -o $(BINARY) .

test:
	go test ./...

e2e: build
	E2E_BINARY=$(BINARY) go -C ../e2e test -v ./...

clean:
	rm -rf bin
