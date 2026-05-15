# CLI Task Tracker — Requirements

A command-line application for managing a personal todo list. Users add tasks, list them, mark them complete, and delete them. The list persists across invocations so the same data is available the next time the program runs.

The application is invoked from a shell as a single, short-lived process — no daemon, no background service. Every operation goes through a full load → mutate → save cycle on disk.

## Functional Requirements

### FR-1: Add a task
The user provides a title and the system records a new task. Each new task is assigned a unique identifier, captured with its creation timestamp, and starts in the *active* (not done) state. After creation, the system prints the new task's identifier so the user can reference it in subsequent commands.

### FR-2: List tasks
The user requests the current task list and the system prints one task per line. Each line shows the identifier, a status marker indicating whether the task is done, and the title.

### FR-3: Filter the listing by status
The user can restrict the listing to a subset by status: `all`, `active`, or `done`. The default is `all`. The filter is supplied as a flag on the list command (e.g. `list --status active`).

### FR-4: Mark a task complete
The user provides a task identifier and the system flips that task's done flag. The change is persisted immediately.

### FR-5: Delete a task
The user provides a task identifier and the system removes that task from the list. The change is persisted immediately and the identifier is not reused for any future task.

### FR-6: Report invalid input
If the user references an identifier that does not exist (on `complete` or `delete`), or supplies an unknown subcommand or flag, the system prints a human-readable error message and exits with a non-zero status.

## Technical Requirements

### TR-1: Persistent state file
State is stored in a single JSON file on disk (default: `tasks.json` in the current working directory). On startup, the absence of this file is *not* an error — it is treated as an empty store. This allows a first-run user to invoke any command without seeding the file by hand. Any other I/O failure (permissions, unreadable file) is surfaced as an error.

### TR-2: Read–modify–write cycle
Every mutating command (`add`, `complete`, `delete`) reads the entire state file, mutates the in-memory representation, and writes it back. No long-lived process holds state between invocations — each invocation is a cold start. This is the simplest correct model for a single-user CLI.

### TR-3: Atomic save
Saves are atomic. The system writes the new state to a temporary file in the same directory as the state file, then renames it over the real path. A crash mid-write cannot leave a half-written state file, because rename within the same filesystem is an atomic operation. The user always sees either the previous contents or the new contents — never a torn write.

### TR-4: Task data model
Each task carries at minimum:
- a unique identifier
- a title (free-form text)
- a *done* flag
- a creation timestamp

Identifiers are unique within the file and remain stable for the lifetime of a task. Deleting a task does not reuse its identifier.

### TR-5: Argument parsing
The first positional argument selects the subcommand (`add`, `list`, `complete`, `delete`). Subsequent positional arguments and flags are parsed per-subcommand. Verify commands in this document use `<run>` as a placeholder for the language-specific run command (e.g. `go run .`, `cargo run --`).

### TR-6: Corrupt state file handling
If the state file exists but its JSON is malformed, the system prints a clear error and exits non-zero. It *must not* silently overwrite the corrupt file with an empty store — doing so would erase the user's data the moment they ran any command. The user must be allowed to inspect or back up the file before deciding what to do.

### TR-7: Exit codes
Exit codes are meaningful so the program composes cleanly in shell pipelines (e.g. `cmd1 && cmd2`):
- `0` — success
- `1` — user error (unknown identifier, invalid flag, missing argument)
- `2` — internal or unexpected error (state file unreadable, malformed JSON)

### TR-8: Scale
The system must remain responsive — no perceptible delay on any command — with at least 1000 tasks in the file. Rewriting the entire file on every mutation is acceptable at this scale; correctness is the priority, not throughput. The system must not contain any O(n²) operation that would degrade past this point.

### TR-9: Stable list ordering
The order in which tasks appear in `list` is stable across invocations. Deletion preserves the order of the remaining tasks (shift the tail down rather than swap-remove). Swap-remove — moving the last element into the hole left by the deleted one — is faster but reorders the visible output, which is undesirable for a user-facing list.

## Acceptance Criteria

The application is considered complete when all of the following hold:

- **AC-1 — Persistence across runs:** `<run> add "buy milk"` followed by `<run> add "write roadmap"`, then closing the terminal and running `<run> list` in a fresh session, prints both tasks.
- **AC-2 — Status filtering:** After adding a task and running `<run> complete <id>`, `<run> list --status done` shows it and `<run> list --status active` does not.
- **AC-3 — Deletion:** After `<run> delete <id>`, `<run> list` no longer shows that task.
- **AC-4 — Scale:** `for i in $(seq 1 1000); do <run> add "task $i"; done` completes without error and `<run> list | wc -l` reports 1000.
- **AC-5 — Corrupt-file safety:** `echo garbage > tasks.json` followed by any command exits non-zero with a readable error, and `tasks.json` is left untouched.
- **AC-6 — Invalid identifier:** `<run> complete 999999` (or any unused id) prints a clear error and exits with code `1`.
