# CLI Task Tracker

A command-line todo list: add tasks, list them, mark them complete, delete them. Each run is a short-lived process that loads the list from disk, applies one change, and saves it back — the data survives across runs.

## Functional Requirements

### FR-1: Add a task

`add "<title>"` records a new task with a unique identifier, a creation timestamp, and an *active* (not done) state, then prints just the identifier on stdout so the next command — or a script — can capture it. Identifiers stay stable for a task's lifetime and are never reused, even after deletion.

### FR-2: List tasks

`list` prints one task per line: identifier, a done/active marker, and the title. The order is stable across runs, and deleting a task keeps the remaining ones in their existing order. `--status all|active|done` filters the listing; the default is `all`.

### FR-3: Complete and delete

`complete <id>` marks the task done; `delete <id>` removes it. Both persist the change immediately.

### FR-4: Persistence that can't be torn

State lives in a single JSON file, `tasks.json` in the working directory. A missing file means an empty list, not an error — a first run needs no setup. Saves are atomic: a crash mid-save can never leave a half-written file; it always holds either the old contents or the new.

### FR-5: Errors are loud and safe

An unknown identifier, subcommand, or flag prints a clear message and exits with code `1`. A state file that exists but can't be read or parsed prints a clear message and exits with code `2` — and the file is left untouched, never overwritten with an empty list. Exit `0` means success.

### FR-6: Stays quick at a thousand tasks

Every command responds without perceptible delay with at least 1000 tasks in the file.

## Acceptance Criteria

- AC-1: `<run> add "buy milk"` then `<run> add "write roadmap"`, then `<run> list` in a fresh terminal session, prints both tasks.
- AC-2: After `<run> complete <id>`, `<run> list --status done` shows the task and `<run> list --status active` does not.
- AC-3: After `<run> delete <id>`, `<run> list` no longer shows that task.
- AC-4: `for i in $(seq 1 1000); do <run> add "task $i"; done` completes without error and `<run> list | wc -l` reports 1000.
- AC-5: `echo garbage > tasks.json` followed by any command exits with code `2` and a readable error, and `tasks.json` is left untouched.
- AC-6: `<run> complete 999999` (or any unused id) prints a clear error and exits with code `1`.
