# Learning by Building

A personal, project-based curriculum (CLI tools → advanced systems). The point is to internalize concepts that are usually taken for granted by building them by hand and to deeply understand the tools being used (programming language, libraries, etc).

## Repo layout

Monorepo. Each project lives in its own top-level folder named `<level>-<slug>` (e.g. `20-cli-task-tracker`). Folders are created as projects are started and stay in the repo once they meet the project's acceptance criteria — no separate "completed" state.

The curriculum itself lives in [curriculum.yaml](curriculum.yaml): 100 projects ordered by `level` (20 → 2000). The level number is purely an ordering — gaps aren't difficulty units, just room to insert later.

Each project entry has: `description`, `concepts`, `requirements` (acceptance criteria), `recommended_languages`. A project is done when every requirement ticks.

## Per-project ROADMAP.md

Every project folder contains a `ROADMAP.md`. It tracks the *actual* steps user needs to build the project. If not present, we need to create it.

Format:

- Short title at the top, then `## Phase N — <name>` sections in build order.
- Each phase is a flat list of `- [ ]` steps, ending with a `- [ ] **Verify:** ...` checkbox naming a concrete command or observation user can run end-to-end so he can report pass/fail.
- Phases slice end-to-end (a user-visible behavior shipped per phase), not horizontally (don't do "all data models" then "all I/O" — there's nothing to demo at the end).
- For any step whose meaning isn't obvious from the bullet alone (e.g. "atomic rename," "swap-remove," "exit code 2," "idempotent"), follow it with a short indented explainer paragraph — 1-3 sentences covering *what it means and why*. Skip the explainer on self-evident steps so the file doesn't get noisy.
- Don't restate anything already in `curriculum.yaml` (description, concepts, requirements, recommended languages). The ROADMAP is the actual build steps, not a copy of the spec.

## Default language

Go. Always use the very last version of the language and retrieve the corresponding documentation.

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

In practice: when user is mid-project, guide him towards the next steps. First provide some debug statements to better understand the structures that we are dealing with. Encourage the user to execute the program with just the debug statements in place. Then once the structures are better understood, help the user get to the next step. Provide the necessary syntax to complete the immediate next step, be very clear about where the changes need to be done. Keep the debug statements in syntax until the user decides to manually delete it.

## Curriculum convention worth knowing

Each level deliberately bans the abstractions that *later* levels will teach you to build. If a project description says "use the stdlib's X," that's load-bearing — a later level exists to make you build X by hand. Don't suggest hand-rolling something the project says to lean on, and don't suggest importing a library for something the project says to build.
