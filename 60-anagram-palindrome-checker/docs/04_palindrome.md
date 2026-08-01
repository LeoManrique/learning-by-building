# 04 — Palindrome

**Goal:** FR-1 working end to end: verdict line on stdout, exit 0 on match, 1 on no-match.

- A string is a palindrome when the normal form reads the same from both ends. Two ways in: walk two indexes from the ends toward the middle, or build the reversed form and compare. Pick one, but understand both.
- Check your solution against the empty string — it should say yes *without* a special case. Work out why.
- Wire the verdict: one human-readable line, then exit with the right code.

**Done when:**

- [ ] AC-1 passes (the panama string is a palindrome, exit 0).
- [ ] AC-2 passes (the empty string is a palindrome, exit 0).
- [ ] AC-3 passes (`hello` is not, exit 1).
