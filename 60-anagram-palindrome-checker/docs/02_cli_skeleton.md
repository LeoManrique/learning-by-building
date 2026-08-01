# 02 — CLI skeleton

**Goal:** the program knows which check was asked for and rejects everything else with exit code 2 (FR-3, usage half).

- Start by seeing what the program actually receives: log the raw argument list at the top of main and run with different argument shapes. Note what the first element is.
- Dispatch on the subcommand: `palindrome` needs exactly one argument after it, `anagram` exactly two.
- Everything else — no arguments, unknown subcommand, wrong count — prints a usage line and exits 2. Look up how your language sets the process exit code, and why usage errors go to **stderr** while verdicts will go to **stdout**.
- The two checks themselves can just print a placeholder for now.

**Done when:**

- [ ] The AC-8 test passes (no arguments, unknown subcommand, or wrong count → usage line, exit 2); the other seven still fail.
