BINARY := $(abspath ../../target/debug/cli-task-tracker-rust)

.PHONY: build test e2e clean

build:
	cargo build

test:
	cargo test

e2e: build
	E2E_BINARY=$(BINARY) go -C ../e2e test -v ./...

clean:
	cargo clean --package cli-task-tracker-rust
