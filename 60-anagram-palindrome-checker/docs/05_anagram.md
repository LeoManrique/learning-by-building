# 05 — Anagram

**Goal:** FR-2 working, whole suite green.

- Sort-based equivalence: two strings are anagrams when their sorted normal forms are equal. Turn each normal form into a sequence of characters, sort both, compare for equality.
- Log both sorted forms side by side for `"listen"` / `"Silent"` before comparing — seeing the two identical sequences is the whole idea in one line of output.
- AC-7 (swapping arguments never changes the answer) should pass without any extra code. Work out why the approach guarantees it.
- Verdict line and exit codes work exactly like the palindrome check.

**Done when:**

- [ ] AC-4 passes (`listen` / `Silent` are anagrams, exit 0).
- [ ] AC-5 passes (the multi-word pair is, exit 0).
- [ ] AC-6 passes (`hello` / `world` are not, exit 1).
- [ ] AC-7 passes (swapped arguments give the same answer).
- [ ] `just e2e` shows all 8 tests passing — project done.
