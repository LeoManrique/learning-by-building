# CLI Task Tracker — Roadmap

## Phase 1 — Persistent add + list

- [ ] Decide on a state file path (e.g. `tasks.json` in the working directory).
- [ ] Define a `Task` struct with fields for id, title, done flag, creation timestamp.
- [ ] Sketch a `Store` type holding `[]Task` plus a path; give it `Load()` and `Save()` methods.
- [ ] In `Load`, treat a missing file as an empty store (don't error).
    A first-run user shouldn't have to seed the file by hand. `errors.Is(err, os.ErrNotExist)` after `os.ReadFile` is the canonical check.
- [ ] In `Save`, write to a temp file in the same directory, then `os.Rename` over the real path.
    Atomic rename: a crash mid-write can't leave a half-written `tasks.json`, because rename within the same filesystem is atomic. Temp-file + rename is the standard pattern for "either the new contents or the old contents, never a torn write."
- [ ] Parse `os.Args` to dispatch on the first positional arg: `add` and `list` for now.
- [ ] Implement `add <title>`: load → append task with a fresh id → save → print the new id.
    Read-modify-write: every mutating command does the full Load/mutate/Save cycle. There is no long-lived process holding state — each invocation is a cold start.
- [ ] Implement `list`: load → print one line per task (id, status marker, title).
- [ ] **Verify:** `go run . add "buy milk"`, then `go run . add "write roadmap"`, then `go run . list` prints both tasks; closing the terminal and re-running `go run . list` still prints them.

## Phase 2 — Complete, delete, and status filter

- [ ] Add `complete <id>`: load → find task by id → set done true → save.
- [ ] Add `delete <id>`: load → remove the matching task → save.
    Slice removal in Go: `s = append(s[:i], s[i+1:]...)` preserves order; `s[i] = s[len(s)-1]; s = s[:len(s)-1]` (swap-remove) is O(1) but reorders. List output should stay stable, so prefer the first form here.
- [ ] Add a `--status all|active|done` flag on `list` (default `all`).
- [ ] On unknown id for complete/delete, print a clear error and exit non-zero.
    Exit codes matter for shell composition (`cmd1 && cmd2`). Use `0` for success and `1` for "user asked for something invalid" (unknown id, bad flag, missing argument).
- [ ] **Verify:** add a task, `complete <id>`, then `list --status done` shows it and `list --status active` doesn't; `delete <id>` then `list` no longer shows it.

## Phase 3 — Robustness and scale

- [ ] On `Load`, if the file exists but JSON is malformed, print a clear error and exit non-zero — do *not* silently overwrite with an empty store.
    Silently resetting a corrupted file would erase the user's data the moment they ran any command. Better to surface the problem so they can inspect or back up the file before deciding what to do.
- [ ] In exit codes, distinguish user errors (`1`) from unexpected/internal errors like "file unreadable" or "malformed JSON" (`2`).
- [ ] Generate 1000 tasks and confirm `add`, `list`, `complete`, `delete` all stay responsive.
    A naive read-modify-write that rewrites the whole file each mutation is fine at this scale — correctness is the requirement, not throughput. Just sanity-check that nothing is accidentally O(n²).
- [ ] **Verify:** `for i in $(seq 1 1000); do go run . add "task $i"; done` completes; `go run . list | wc -l` reports 1000; `echo garbage > tasks.json` followed by any command exits non-zero with a readable error.
