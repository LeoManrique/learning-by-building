// Black-box acceptance tests for the anagram & palindrome checker.
//
// The suite never imports the implementation: it executes the compiled
// binary named by the E2E_BINARY environment variable and asserts on exit
// codes. Run it via `make e2e` inside an implementation folder.
package e2e

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// binary is the path of the program under test, taken from E2E_BINARY.
var binary string

func TestMain(m *testing.M) {
	binary = os.Getenv("E2E_BINARY")
	if binary == "" {
		fmt.Fprintln(os.Stderr, "e2e: E2E_BINARY is not set; run `make e2e` from an implementation folder")
		os.Exit(2)
	}
	os.Exit(m.Run())
}

// run executes the binary with args and returns its exit code and combined
// stdout+stderr output.
func run(t *testing.T, args ...string) (code int, output string) {
	t.Helper()
	out, err := exec.Command(binary, args...).CombinedOutput()
	if err == nil {
		return 0, string(out)
	}
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		t.Fatalf("could not run %q: %v", binary, err)
	}
	return exitErr.ExitCode(), string(out)
}

// expectExit runs the binary and fails the test if the exit code differs.
func expectExit(t *testing.T, want int, args ...string) {
	t.Helper()
	code, out := run(t, args...)
	if code != want {
		t.Errorf("exit code = %d, want %d\nargs: %q\noutput:\n%s", code, want, args, out)
	}
}

// AC-1: the classic phrase is a palindrome.
func TestPalindromeClassicPhrase(t *testing.T) {
	expectExit(t, 0, "palindrome", "A man a plan a canal Panama")
}

// AC-2: the empty string is a palindrome.
func TestPalindromeEmptyString(t *testing.T) {
	expectExit(t, 0, "palindrome", "")
}

// AC-3: a plain word that isn't a palindrome exits 1.
func TestPalindromePlainWordIsNotOne(t *testing.T) {
	expectExit(t, 1, "palindrome", "hello")
}

// AC-4: anagram check ignores case.
func TestAnagramListenVsSilent(t *testing.T) {
	expectExit(t, 0, "anagram", "listen", "Silent")
}

// AC-5: anagram check works on multi-word phrases and ignores spaces.
func TestAnagramMultiWordPhrase(t *testing.T) {
	expectExit(t, 0, "anagram", "conversation", "voices rant on")
}

// AC-6: unrelated words are not anagrams; exits 1.
func TestAnagramUnrelatedWords(t *testing.T) {
	expectExit(t, 1, "anagram", "hello", "world")
}

// AC-7: swapping the two anagram arguments must not change the answer.
func TestAnagramOrderIndependence(t *testing.T) {
	ab, _ := run(t, "anagram", "a", "b")
	ba, _ := run(t, "anagram", "b", "a")
	if ab != ba {
		t.Errorf("anagram a b exited %d but anagram b a exited %d", ab, ba)
	}
}

// AC-8a: no arguments is a usage error; exits 2.
func TestUsageNoArguments(t *testing.T) {
	expectExit(t, 2)
}

// AC-8b: an unknown subcommand is a usage error; exits 2.
func TestUsageUnknownSubcommand(t *testing.T) {
	expectExit(t, 2, "nope", "x")
}

// AC-8c: the wrong argument count is a usage error; exits 2.
func TestUsageWrongArgCount(t *testing.T) {
	expectExit(t, 2, "anagram", "only-one")
}
