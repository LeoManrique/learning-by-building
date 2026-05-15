# Learning by Building

A personal, project-based curriculum (CLI tools → advanced systems). The point is to internalize concepts that are usually taken for granted by building them by hand and to deeply understand the tools being used (programming language, libraries, etc).

## Repo layout

Monorepo. Each project lives in its own top-level folder named `<level>-<slug>` (e.g. `20-cli-task-tracker`). Folders are created as projects are started and stay in the repo once they meet the project's acceptance criteria — no separate "completed" state.

The curriculum itself lives in [curriculum.yaml](curriculum.yaml). Projects span 5 phases — Basics, Internals, Systems, Compilers & Languages, Frontier — ordered by `level`. Levels are a build sequence with a soft difficulty trend; gaps are insertion room, not difficulty units.

Each project entry has: `description`, `concepts`, `requirements` (acceptance criteria), `recommended_languages`. A project is done when every requirement ticks.

Default layout: single language at the project root. If user chooses to build a project in both languages (more common in Phase I), use this layout instead:

    <level>-<slug>/
      REQUIREMENTS.md            # language-agnostic, functional + technical requirements
      <slug>-go/                 # Go implementation
      <slug>-rust/               # Rust implementation

## Per-project REQUIREMENTS.md

Every project folder must contain a `REQUIREMENTS.md`. If it's not already created, create it. It describes the project as if a customer's requirements had already been analyzed and translated by a technical analyst — what the software does (functional) and how it has to behave (technical), language-agnostic, before any code is written. If not present, we need to create it.

Format:

- Short title at the top followed by a 1-2 paragraph overview: what the app is, who uses it, and the headline shape of the system (e.g. short-lived CLI process, long-running server, etc.).
- Three sections, in this order:
    - `## Functional Requirements` — what the system does from the user's point of view. Numbered `FR-1`, `FR-2`, … each with a short heading (`### FR-1: Add a task`) and a 1-3 sentence body.
    - `## Technical Requirements` — how the system behaves under the hood: data model, persistence, atomicity, error handling, exit codes, performance targets, ordering guarantees, etc. Numbered `TR-1`, `TR-2`, … same heading + short body shape.
    - `## Acceptance Criteria` — a flat bulleted list of verifiable conditions, numbered `AC-1`, `AC-2`, …. Each one names a concrete command or observation the user can run end-to-end so he can report pass/fail.
- Tone is declarative ("the system records…", "the user provides…"), not imperative ("decide on…", "implement…"). The file describes the finished state, not the build sequence.
- For any requirement whose meaning isn't obvious from the heading alone (e.g. "atomic save," "swap-remove vs preserve-order," "exit code 2 distinguishes user error from internal error"), spend a sentence in the body covering *what it means and why*. Skip the explainer on self-evident requirements so the file doesn't get noisy.
- Language-agnostic. No language-specific syntax, type names, library names, or function names. Acceptance criteria use `<run>` as a placeholder for the language-specific run command (`go run .`, `cargo run --`, etc.).
- Build sequencing (phases, order of work, which feature to ship first) does *not* belong in this file. It's a requirements doc, not a roadmap.

## Programming languages

Go and Rust. Use the latest version of each, with current docs and best practices.

Pick one language per project. Go and Rust are both valid choices. The `recommended_languages` field in `curriculum.yaml` lists which the project leans toward, but the user has free choice. Some projects (especially in Phase I) may end up implemented in both as a deliberate exercise — that's allowed but no longer required.

## Always fix

Do not ignore the errors or warnings when building or running the projects, even if they were already present or not caused by your changes. Do not bypass errors or warnings, actually fix them.

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

## Curriculum convention worth knowing

Each level deliberately bans the abstractions that *later* levels will teach you to build. If a project description says "use the stdlib's X," that's load-bearing — a later level exists to make you build X by hand. Don't suggest hand-rolling something the project says to lean on, and don't suggest importing a library for something the project says to build.
