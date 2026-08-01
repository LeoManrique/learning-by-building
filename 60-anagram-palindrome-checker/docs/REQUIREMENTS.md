# Anagram & Palindrome Checker

A one-shot command-line tool with two checks: is a string a palindrome, and are two strings anagrams of each other. It prints a short verdict and carries the machine-readable answer in the exit code.

## Functional Requirements

### FR-1: Palindrome check

`<run> palindrome "<string>"` answers whether the string reads the same forward and backward. Comparison is case-insensitive and ignores everything that is not a letter, so `"A man a plan a canal Panama"` is a palindrome — and so is the empty string.

### FR-2: Anagram check

`<run> anagram "<string-a>" "<string-b>"` answers whether the two strings contain the same letters, in any order. The same rules apply — case-insensitive, non-letters ignored — so `"listen"` / `"Silent"` and `"conversation"` / `"voices rant on"` are anagrams. Swapping the two arguments never changes the answer.

### FR-3: Exit codes carry the result

Exit `0` means the check passed, `1` means it failed (a real answer of "no"), and `2` means a usage error — missing or unknown subcommand, wrong argument count — where the program never got far enough to answer. A human-readable verdict line is printed alongside; the exit code is for scripts.

## Acceptance Criteria

- AC-1: `<run> palindrome "A man a plan a canal Panama"` prints a match verdict and exits with code `0`.
- AC-2: `<run> palindrome ""` (empty string) prints a match verdict and exits with code `0`.
- AC-3: `<run> palindrome "hello"` prints a no-match verdict and exits with code `1`.
- AC-4: `<run> anagram "listen" "Silent"` prints a match verdict and exits with code `0`.
- AC-5: `<run> anagram "conversation" "voices rant on"` prints a match verdict and exits with code `0`.
- AC-6: `<run> anagram "hello" "world"` prints a no-match verdict and exits with code `1`.
- AC-7: `<run> anagram "a" "b"` and `<run> anagram "b" "a"` produce the same exit code.
- AC-8: `<run>` with no arguments, an unknown subcommand, or the wrong number of arguments prints a usage message and exits with code `2`.
