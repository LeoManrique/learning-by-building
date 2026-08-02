# 01 — Setup

**Goal:** a compiling program and the acceptance suite runnable against it.

- Inside your language subfolder, initialize the project with the toolchain's init command, using the project slug as the name.
- Create the entry point with a main that just prints something.
- Run `just build`, then `just e2e`. All 8 tests should fail — that's the finish line for the project, not a problem to fix now.

**Done when:**

- [ ] `just e2e` builds your binary and reports failures against it.
