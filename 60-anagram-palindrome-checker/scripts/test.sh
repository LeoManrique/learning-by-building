#!/usr/bin/env bash
# Acceptance tests for the anagram & palindrome checker.
# Language-agnostic: pass the program's run command as arguments.
#
#   ./test.sh cargo run --quiet --
#   ./test.sh go run .

set -u

if [[ $# -lt 1 ]]; then
    echo "usage: $0 <run-command...>" >&2
    exit 2
fi

readonly RUN=("$@")

# --- output helpers ----------------------------------------------------------

if [[ -t 1 ]]; then
    readonly GREEN=$'\033[32m' RED=$'\033[31m' RESET=$'\033[0m'
else
    readonly GREEN="" RED="" RESET=""
fi

pass_count=0
fail_count=0

pass() {
    printf "  %sPASS%s  %s\n" "$GREEN" "$RESET" "$1"
    pass_count=$((pass_count + 1))
}

fail() {
    printf "  %sFAIL%s  %s\n" "$RED" "$RESET" "$1"
    [[ -n "${2:-}" ]] && printf "        %s\n" "$2"
    fail_count=$((fail_count + 1))
}

# --- test primitives ---------------------------------------------------------

# Run the program with the given args; echo only the exit code.
exit_code_of() {
    "${RUN[@]}" "$@" >/dev/null 2>&1
    echo $?
}

# Assert the program exits with $expected when invoked with the trailing args.
expect_exit() {
    local label="$1" expected="$2"
    shift 2

    local actual
    actual=$(exit_code_of "$@")

    if [[ $actual -eq $expected ]]; then
        pass "$label"
    else
        fail "$label" "expected exit $expected, got $actual"
    fi
}

# --- cases -------------------------------------------------------------------

echo "Palindrome:"
expect_exit "AC-1  classic phrase"       0  palindrome "A man a plan a canal Panama"
expect_exit "AC-2  empty string"         0  palindrome ""
expect_exit "AC-3  plain word, not one"  1  palindrome "hello"

echo
echo "Anagram:"
expect_exit "AC-4  listen vs Silent"     0  anagram "listen" "Silent"
expect_exit "AC-5  multi-word phrase"    0  anagram "conversation" "voices rant on"
expect_exit "AC-6  unrelated words"      1  anagram "hello" "world"

# AC-7: argument order must not change the answer.
ab=$(exit_code_of anagram "a" "b")
ba=$(exit_code_of anagram "b" "a")
if [[ $ab -eq $ba ]]; then
    pass "AC-7  order independence"
else
    fail "AC-7  order independence" "a/b=$ab, b/a=$ba"
fi

echo
echo "Usage errors:"
expect_exit "AC-8a no arguments"         2
expect_exit "AC-8b unknown subcommand"   2  nope "x"
expect_exit "AC-8c wrong arg count"      2  anagram "only-one"

# --- summary -----------------------------------------------------------------

echo
echo "$pass_count passed, $fail_count failed"
[[ $fail_count -eq 0 ]]
