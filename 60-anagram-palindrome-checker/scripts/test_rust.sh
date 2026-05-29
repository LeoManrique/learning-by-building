#!/usr/bin/env bash
# Run acceptance tests against the Rust implementation.
set -eu

here="$(cd "$(dirname "$0")" && pwd)"
cd "$here/../anagram-palindrome-checker-rust"
exec "$here/test.sh" cargo run --quiet --
