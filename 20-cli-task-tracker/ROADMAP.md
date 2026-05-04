# CLI Task Tracker — Roadmap

## Phase 1 — Persistent add + list

- [ ] Decide on a state file path (e.g. `tasks.json` in the working directory).
- [ ] Settle on the task fields: id, title, done flag, creation timestamp.
- [ ] On startup, treat a missing state file as an empty store (don't error).
    A first-run user shouldn't have to seed the file by hand. The "file not found" case returns the empty list; everything else propagates as an error.
- [ ] When saving, write to a temp file in the same directory, then atomically rename it over the real path.
    Atomic rename: a crash mid-write can't leave a half-written state file, because rename within the same filesystem is atomic. Temp-file + rename is the standard pattern for "either the new contents or the old contents, never a torn write."
- [ ] Parse the program arguments to dispatch on the first positional arg: `add` and `list` for now.
- [ ] Implement `add <title>`: load → append a task with a fresh id → save → print the new id.
    Read-modify-write: every mutating command does the full load/mutate/save cycle. There is no long-lived process holding state — each invocation is a cold start.
- [ ] Implement `list`: load → print one line per task (id, status marker, title).
- [ ] **Verify:** `<run> add "buy milk"`, then `<run> add "write roadmap"`, then `<run> list` prints both tasks; closing the terminal and re-running `<run> list` still prints them.

## Phase 2 — Complete, delete, and status filter

- [ ] Add `complete <id>`: load → find the task by id → mark it done → save.
- [ ] Add `delete <id>`: load → remove the matching task → save.
    Two ways to remove from a sequence: preserve-order (shift the tail down) or swap-remove (move the last element into the hole). Swap-remove is faster but reorders the output; for stable list output, use preserve-order.
- [ ] Add a `--status all|active|done` flag on `list` (default `all`).
- [ ] On unknown id for `complete`/`delete`, print a clear error and exit non-zero.
    Exit codes matter for shell composition (`cmd1 && cmd2`). Use `0` for success and `1` for "user asked for something invalid" (unknown id, bad flag, missing argument).
- [ ] **Verify:** add a task, `<run> complete <id>`, then `<run> list --status done` shows it and `<run> list --status active` doesn't; `<run> delete <id>` then `<run> list` no longer shows it.

## Phase 3 — Robustness and scale

- [ ] If the state file exists but its JSON is malformed, print a clear error and exit non-zero — do *not* silently overwrite with an empty store.
    Silently resetting a corrupted file would erase the user's data the moment they ran any command. Better to surface the problem so they can inspect or back up the file before deciding what to do.
- [ ] In exit codes, distinguish user errors (`1`) from unexpected/internal errors like "file unreadable" or "malformed JSON" (`2`).
- [ ] Generate 1000 tasks and confirm `add`, `list`, `complete`, `delete` all stay responsive.
    A naive read-modify-write that rewrites the whole file each mutation is fine at this scale — correctness is the requirement, not throughput. Just sanity-check that nothing is accidentally O(n²).
- [ ] **Verify:** `for i in $(seq 1 1000); do <run> add "task $i"; done` completes; `<run> list | wc -l` reports 1000; `echo garbage > tasks.json` followed by any command exits non-zero with a readable error.
