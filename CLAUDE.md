# Learning by Building

A personal, project-based curriculum (CLI tools → advanced systems). The point is to internalize concepts that are usually taken for granted by building them by hand and to deeply understand the tools being used (programming language, libraries, etc).

## Repo layout

Monorepo. Each project lives in its own top-level folder named `<level>-<slug>` (e.g. `20-number-guessing-game`). Folders are created as projects are started and stay in the repo once they meet the project's acceptance criteria — no separate "completed" state.

The curriculum itself lives in [curriculum.yaml](curriculum.yaml). Projects span 5 phases — Basics, Internals, Systems, Compilers & Languages, Frontier — ordered by `level`. Levels are a build sequence with a soft difficulty trend; gaps are insertion room, not difficulty units. The user works through projects sequentially by level.

Each project entry has: `description`, `concepts`, `requirements` (acceptance criteria), `recommended_languages`. A project is done when every requirement ticks.

Project folder layout: `REQUIREMENTS.md` and the numbered step guides live in `docs/`; each language implementation lives in its own subfolder named `<slug>-<language>/`. This applies even when only one language is implemented — the subfolder is always present so a second language can be added later without restructuring. Black-box acceptance tests live in `e2e/` (see "E2E tests"). There is no `scripts/` folder — the e2e suite replaced the old bash harnesses; do not create new ones.

    <level>-<slug>/
      docs/
        REQUIREMENTS.md          # language-agnostic functional requirements + acceptance criteria
        NN_<step>.md             # numbered step guides, one per milestone (see "Per-project step guides")
      e2e/                       # black-box acceptance tests; standalone Go module
      <slug>-go/                 # Go implementation (if present)
        justfile                 # build/test/e2e recipes for this implementation
      <slug>-rust/               # Rust implementation (if present)

The language package/module name inside each subfolder is the slug alone (e.g. `unit-converter`), without the language suffix — the language is already implied by which subfolder the code lives in. In Rust, run `cargo init --name <slug>` inside the subfolder so the crate name in `Cargo.toml` is `<slug>`. In Go, run `go mod init <slug>` inside the subfolder so the module path is `<slug>`. Projects 20 and 140 predate this rule and keep their older `<slug>-<language>` package names; do not retroactively rename them.

## Per-project REQUIREMENTS.md

Every project must contain a `docs/REQUIREMENTS.md`. If it's not already created, create it. It describes what the finished software does, language-agnostic, before any code is written. Keep it short. There is no technical-requirements section: each functional requirement states the observable behavior precisely enough — inputs, outputs, exit codes, what survives a crash or restart, what happens on bad input — that the technical decisions can be inferred from it.

Format:

- Short title followed by a 1-2 sentence overview: what the app is and its headline shape (short-lived CLI process, long-running server, etc.).
- Two sections, in this order:
    - `## Functional Requirements` — numbered `FR-1`, `FR-2`, … each with a short heading (`### FR-1: Add a task`) and a 1-3 sentence body. Observable behavior only; internal mechanics appear only through their visible effects (say "a crash mid-save never leaves a half-written file", not "write to a temp file and rename").
    - `## Acceptance Criteria` — a flat bulleted list of verifiable conditions, numbered `AC-1`, `AC-2`, …. Each one names a concrete command or observation that both the user and the e2e suite can check end-to-end: exit codes, printed output, files left on disk. Avoid criteria only a human can judge — the e2e suite must be able to assert every one.
- Tone is declarative ("the system records…"), in plain, everyday language — no academic or jargony phrasings. The file describes the finished state, not the build sequence; no roadmap or phases.
- Language-agnostic. No language-specific syntax, type names, library names, or function names. Acceptance criteria use `<run>` as a placeholder for the language-specific run command (`go run .`, `cargo run --`, etc.).
- List only the minimum requirements needed to exercise the project's headline concepts. Do not pad the spec with extra features, configurability knobs, robustness guarantees, or edge-case handling unless the curriculum entry specifically calls for them — if the user wants more, he adds those requirements himself.

## Per-project step guides

Alongside `REQUIREMENTS.md`, `docs/` holds a short sequence of numbered guides — `01_<step>.md`, `02_<step>.md`, … (lowercase snake_case) — one per milestone, created when the project starts. They guide *what* to do, not *how*: concise, high level, never an exact step-by-step, never code. Writing them is Claude's job; project 60 is the reference example.

Each guide has a fixed shape:

- `# NN — Title`, then a one-line **Goal:** tied to the functional requirement(s) the milestone serves.
- A few bullets of what to do. Name concepts so they can be looked up, but leave the composition to the user; where a standard-library facility is the intended tool, say it exists ("your standard library has a way to…") without naming it. At least one bullet is a logging step (see "Logging as visualization") — run the program and see the structure before building on it.
- A closing **Done when:** checklist — `- [ ]` items naming the acceptance tests that pass at the end of the milestone, each with a few words of what the test shows so the box is meaningful without opening `REQUIREMENTS.md`. The user ticks the boxes as he goes; the boxes are how the agent locates current progress. Progress is measured by `just e2e`, never by feel.

Rules:

- Language-agnostic, same rule as `REQUIREMENTS.md`: no language names, no type, library, or function names, no language-specific syntax.
- Ordering carries the pedagogy: 01 is always setup and ends with a first `just e2e` run where every test fails (the finish line made visible); shared foundations get their own milestone before the features that stand on them; the last milestone ends with the whole suite green.
- Flag what's worth understanding ("work out why the empty string passes without a special case") — never explain it away in the guide.
- Guides reference FR and AC numbers, so whoever changes `REQUIREMENTS.md` updates the guides in the same change — same no-drift rule as the e2e suite.
- During `help` and `review`, the current milestone is the first guide with an unticked box, and its guide anchors what the next step is. `review` cross-checks the ticks against the `just e2e` results — a ticked box the suite disputes gets flagged, and vice versa — but the user's boxes are his to tick; the agent never edits them.

## E2E tests

Every project has an `e2e/` folder: a standalone Go module named `<slug>-e2e` holding black-box acceptance tests — one plain test function per acceptance criterion, run sequentially, with the AC number in a comment above it. No table-driven tests: the user prefers separate, readable functions. Writing and maintaining this suite is Claude's job (see "How user uses AI in this repo"); it replaces the old `scripts/*.sh` harnesses.

- The acceptance criteria in `docs/REQUIREMENTS.md` are the single source of truth, and the suite mirrors them one-to-one: every AC has a test function carrying its number, every test function traces back to an AC, and no test asserts behavior the ACs don't state. An AC that bundles several concrete shapes may map to one function per shape (`AC-8a`, `AC-8b`, …). Whoever changes the ACs updates the suite in the same change — the two are never allowed to drift.
- The suite never imports or builds the implementation. It executes the compiled binary whose path arrives in the `E2E_BINARY` environment variable and asserts on exit codes and output. That is what keeps it implementation-language-agnostic: a Go and a Rust implementation are exercised by the same suite, each pointing `E2E_BINARY` at its own binary.
- The suite is always written in Go regardless of the implementation language, and uses only the standard library (`testing`, `os/exec`). One file unless it genuinely outgrows that; shared setup goes in small helpers (`run`, `expectExit`), not tables.
- Each implementation folder carries a `justfile` with at least `build`, `test` (unit), and `e2e` recipes. `just e2e` compiles the implementation and runs the shared suite against it (`E2E_BINARY=<binary> go -C ../e2e test ./...`).
- Unit tests inside the implementation are the user's job (core logic, stdlib `testing`, separate functions over tables). From the API-level projects (level 840 up), testify may be adopted to mirror the job stack.

## Programming languages

Go and Rust. Use the latest version of each, with current docs and best practices.

**Go-first, Rust paused (since Aug 2026):** the user starts a Go-stack job on 2026-11-01, so all new projects are implemented in Go and no new Rust implementations are started until further notice. Finished Rust work stays in the repo; the `<slug>-rust/` folder convention stays for when Rust resumes.

Pick one language per project. The `recommended_languages` field in `curriculum.yaml` lists which the project leans toward, but the user has free choice.

## Always fix

Do not ignore the errors or warnings when building or running the projects, even if they were already present or not caused by your changes. Do not bypass errors or warnings, actually fix them.

## Always modern

All code suggestions — in any command or mid-project assistance — must use the most modern idioms and syntax for the language version declared in the project. Check the project's manifest (edition, language version, toolchain file) before suggesting code, and prefer current-version forms over older equivalents even when both compile.

Match conventions already present in the file before introducing new ones. If the existing code uses a modern form, do not regress it in your suggestions.

## Always explain

If user asks for help, don't just help give him instructions to complete the project, but provide explanations that will create a deep understanding of the topic.

## Logging as visualization

Suggest the user to log how the system works, and provide cases to change what the log outputs so that the behavior of the technology is better understood and internalized.

## How user uses AI in this repo

**Use Claude for:**
- Explaining concepts in depth — go past the one-liner. Cover what the thing *is*, why it exists, the tradeoffs, common pitfalls, and how it connects to ideas user already knows. Use concrete examples; assume he wants the mental model, not just the definition.
- Explain which tools (libraries, frameworks, packages, etc...) to use for the required task and how do they work.
- Suggesting best practices — and explain *why* each is best practice (what it prevents, where it came from, when it doesn't apply).
- Debug statements to better understand the structures
- Help with latest version of Go's Syntax (or whichever language the current project uses)
- Checking work against a project's acceptance criteria
- Repetitive / boilerplate code
- Writing, running, and reviewing E2E test output

**Do NOT use Claude for:**
- Writing most of the core logic of a project — that's the whole point of the exercise
- Actions the user can do themselves with a single or few commands. Output the commands for the user to run instead of performing them. Examples: project initialization (`go mod init`, `cargo init`, etc.), scaffolding from a generator, adding a dependency, creating an empty file. The user learns the toolchain by using it, and typing one command is faster than waiting for the AI to write the files it produces.

In practice: when user is mid-project, guide him towards the next steps. First provide some debug statements to better understand the structures that we are dealing with. Encourage the user to execute the program with just the debug statements in place. Then once the structures are better understood, help the user get to the next step. Provide the necessary syntax to complete the immediate next step, be very clear about where the changes need to be done. Keep the debug statements in syntax until the user decides to manually delete it.

Pacing: introduce one new concept or syntax form at a time including debug statements. A code block that implements user-visible behavior end-to-end means you've written the project for them.

## Commands

These commands are the only ways to invoke help in this repo. If a message does not begin with one of these commands, refuse to act and reply with the list of available commands. The rule applies to messages that initiate a new request — once a command is active, follow-up exchanges within that thread continue normally without needing the keyword again.

### `start <level> [language]`

Project initialization. The level must match a `level` in `curriculum.yaml`. The language defaults to `go` when omitted; the only other accepted value is `rust` (currently paused — point that out if requested, but obey).

Behavior:

- If the project folder `<level>-<slug>` does not exist, create it (slug is derived from the matching `id` in `curriculum.yaml`).
- Create `docs/REQUIREMENTS.md` if not present, following the format in "Per-project REQUIREMENTS.md". If it is already present, leave it alone.
- Create the numbered step guides in `docs/` if not present, per "Per-project step guides". If any are present, leave them alone.
- Create `e2e/` if not present: a standalone Go module `<slug>-e2e` with one black-box test per acceptance criterion, per "E2E tests".
- Create the language subfolder `<slug>-<language>/` with its `justfile` if not present.
- Print the toolchain command the user runs inside the language subfolder to initialize the project: `go mod init <slug>` for Go, `cargo init --name <slug>` for Rust. Do not run the init yourself.

### `help`

The user is stuck and wants to be unstuck without being handed the answer. Provide hints, point at the relevant concept, and lead toward a next step the user can take. If the structures or runtime behavior are unclear, suggest log statements to add and ask the user to run the program with them. Never produce a code block that implements user-visible behavior end-to-end.

### `review`

Check current progress against `REQUIREMENTS.md` and code quality. Assume the user has made some progress — likely a step or a few — not that the project is finished. The goal is to confirm what's on the right track and nudge toward the next step, not to audit the whole spec.

- If the implementation builds, run `just e2e` and read the results as the acceptance-criteria check instead of eyeballing the code. Briefly mention which acceptance criteria are met and which are not. Don't mention details of steps that are not the immediate next one.
- Flag code-quality issues (idioms, naming, error handling) appropriate to the project's level — only when they are actually wrong or worth changing. If you would caveat a point with "not wrong, just…" or "fine for now," omit it entirely. Silence is the correct output when there is nothing to flag. Do not push abstractions that later levels are designed to teach (see "Curriculum convention worth knowing").
- End with the immediate next step — one concept or one syntax form at a time, per the pacing rule.

### `explain <topic or tool>`

Deep explainer for a concept, library, framework, or language feature. Cover what it is, why it exists, the main tradeoffs, common pitfalls, and how it connects to ideas already in the curriculum. Use concrete examples; assume the user wants the mental model, not just the definition.

### `curriculum`

Repo maintenance mode: discuss and apply changes to `curriculum.yaml`, `CLAUDE.md`, `README.md`, or repo-wide conventions. This is normal engineering collaboration — the learning guardrails (hints-only, no end-to-end code) apply to project code, and none is written here.

## Curriculum convention worth knowing

Each level deliberately bans the abstractions that *later* levels will teach you to build. If a project description says "use the stdlib's X," that's load-bearing — a later level exists to make you build X by hand. Don't suggest hand-rolling something the project says to lean on, and don't suggest importing a library for something the project says to build.
