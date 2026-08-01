// Black-box acceptance tests for the CLI task tracker.
//
// The suite never imports the implementation: it executes the compiled
// binary named by the E2E_BINARY environment variable and asserts on exit
// codes, output, and the state file left on disk. Run it via `make e2e`
// inside an implementation folder.
//
// Every test gets its own temporary working directory, because the tracker
// keeps its state in tasks.json in the current working directory.
package e2e

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

// runIn executes the binary inside dir and returns its exit code, stdout,
// and stderr.
func runIn(t *testing.T, dir string, args ...string) (code int, stdout, stderr string) {
	t.Helper()
	cmd := exec.Command(binary, args...)
	cmd.Dir = dir
	var outBuf, errBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			t.Fatalf("could not run %q: %v", binary, err)
		}
		code = exitErr.ExitCode()
	}
	return code, outBuf.String(), errBuf.String()
}

// addTask adds a task and returns the identifier the program printed.
func addTask(t *testing.T, dir, title string) string {
	t.Helper()
	code, stdout, stderr := runIn(t, dir, "add", title)
	if code != 0 {
		t.Fatalf("add %q exited %d, want 0\nstderr:\n%s", title, code, stderr)
	}
	id := strings.TrimSpace(stdout)
	if id == "" {
		t.Fatalf("add %q printed nothing; expected the new task's identifier on stdout", title)
	}
	return id
}

// listTasks runs `list` with the extra args and returns stdout, failing the
// test if the command does not exit 0.
func listTasks(t *testing.T, dir string, extra ...string) string {
	t.Helper()
	args := append([]string{"list"}, extra...)
	code, stdout, stderr := runIn(t, dir, args...)
	if code != 0 {
		t.Fatalf("%v exited %d, want 0\nstderr:\n%s", args, code, stderr)
	}
	return stdout
}

// nonEmptyLines counts the non-empty lines in s — the `list | wc -l` view.
func nonEmptyLines(s string) int {
	n := 0
	for line := range strings.Lines(s) {
		if strings.TrimSpace(line) != "" {
			n++
		}
	}
	return n
}

// AC-1: two adds followed by a fresh `list` print both tasks.
func TestPersistenceAcrossRuns(t *testing.T) {
	dir := t.TempDir()
	addTask(t, dir, "buy milk")
	addTask(t, dir, "write roadmap")
	out := listTasks(t, dir)
	if !strings.Contains(out, "buy milk") || !strings.Contains(out, "write roadmap") {
		t.Errorf("list is missing a task added in an earlier run:\n%s", out)
	}
}

// AC-2: a completed task shows under --status done and not under --status active.
func TestStatusFiltering(t *testing.T) {
	dir := t.TempDir()
	id := addTask(t, dir, "buy milk")
	if code, _, stderr := runIn(t, dir, "complete", id); code != 0 {
		t.Fatalf("complete %s exited %d, want 0\nstderr:\n%s", id, code, stderr)
	}
	if done := listTasks(t, dir, "--status", "done"); !strings.Contains(done, "buy milk") {
		t.Errorf("list --status done does not show the completed task:\n%s", done)
	}
	if active := listTasks(t, dir, "--status", "active"); strings.Contains(active, "buy milk") {
		t.Errorf("list --status active still shows the completed task:\n%s", active)
	}
}

// AC-3: a deleted task no longer appears in `list`.
func TestDeletion(t *testing.T) {
	dir := t.TempDir()
	id := addTask(t, dir, "buy milk")
	if out := listTasks(t, dir); !strings.Contains(out, "buy milk") {
		t.Fatalf("list does not show the task before deletion:\n%s", out)
	}
	if code, _, stderr := runIn(t, dir, "delete", id); code != 0 {
		t.Fatalf("delete %s exited %d, want 0\nstderr:\n%s", id, code, stderr)
	}
	if out := listTasks(t, dir); strings.Contains(out, "buy milk") {
		t.Errorf("list still shows the deleted task:\n%s", out)
	}
}

// AC-4: 1000 adds complete without error and `list` prints 1000 lines.
func TestScaleThousandTasks(t *testing.T) {
	dir := t.TempDir()
	for i := 1; i <= 1000; i++ {
		addTask(t, dir, fmt.Sprintf("task %d", i))
	}
	if got := nonEmptyLines(listTasks(t, dir)); got != 1000 {
		t.Errorf("list prints %d lines, want 1000", got)
	}
}

// AC-5: a corrupt state file makes any command exit 2 with a readable error,
// and the file is left untouched.
func TestCorruptFileSafety(t *testing.T) {
	dir := t.TempDir()
	stateFile := filepath.Join(dir, "tasks.json")
	if err := os.WriteFile(stateFile, []byte("garbage\n"), 0o644); err != nil {
		t.Fatalf("could not seed corrupt state file: %v", err)
	}
	code, _, stderr := runIn(t, dir, "list")
	if code != 2 {
		t.Errorf("exit code = %d, want 2 for a corrupt state file", code)
	}
	if strings.TrimSpace(stderr) == "" {
		t.Errorf("expected a readable error on stderr, got nothing")
	}
	after, err := os.ReadFile(stateFile)
	if err != nil {
		t.Fatalf("could not re-read state file: %v", err)
	}
	if string(after) != "garbage\n" {
		t.Errorf("corrupt state file was modified; the user must be able to inspect it\nnow contains:\n%s", after)
	}
}

// AC-6: completing an identifier that does not exist exits 1 with a clear error.
func TestInvalidIdentifier(t *testing.T) {
	dir := t.TempDir()
	addTask(t, dir, "buy milk")
	code, _, stderr := runIn(t, dir, "complete", "999999")
	if code != 1 {
		t.Errorf("exit code = %d, want 1 for an unknown identifier", code)
	}
	if strings.TrimSpace(stderr) == "" {
		t.Errorf("expected a clear error on stderr, got nothing")
	}
}
