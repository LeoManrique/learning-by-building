# Anagram & Palindrome Checker

A small command-line program that runs two checks on text and reports the result with an exit code. The first check answers "is this string a palindrome?" — does it read the same forward and backward, ignoring case, spaces, and punctuation. The second check answers "are these two strings anagrams of each other?" — do they contain the same letters, regardless of order or spacing. The program runs once per invocation, prints its verdict, and exits.

## Functional Requirements

### FR-1: Palindrome check

The user can ask whether a single string is a palindrome. Letters are compared case-insensitively, and anything that is not a letter (spaces, punctuation, digits) is ignored when checking. So `"A man a plan a canal Panama"` is a palindrome, and the empty string is a palindrome.

### FR-2: Anagram check

The user can ask whether two strings are anagrams of each other. As with the palindrome check, comparison is case-insensitive and non-letters are ignored. So `"listen"` and `"Silent"` are anagrams, and so are `"conversation"` and `"voices rant on"`.

### FR-3: Command shape

The program is invoked with a subcommand that selects the check, followed by the input string(s). The shapes are:

- `<run> palindrome "<string>"`
- `<run> anagram "<string-a>" "<string-b>"`

Any other shape (missing subcommand, wrong number of arguments, unknown subcommand) is a usage error.

### FR-4: Output

On a successful check the program prints a short, human-readable verdict — for example `match` or `no match`. The exit code (see TR-2) carries the machine-readable answer; the printed line is for the person watching the terminal.

## Technical Requirements

### TR-1: Normalization before comparison

Both checks share the same normalization step: fold to a single case, then drop everything that is not a letter. Comparison happens on the normalized form, never on the raw input. This is what lets `"A man a plan a canal Panama"` be a palindrome and `"conversation"` / `"voices rant on"` be anagrams.

### TR-2: Exit codes carry the result

The exit code is the primary signal so the program is usable from scripts. `0` means the check passed (palindrome / anagram). `1` means the check failed (not a palindrome / not an anagram). `2` is reserved for usage errors — wrong subcommand, wrong argument count, unknown flag. The distinction between `1` and `2` matters: `1` is a real, well-formed answer of "no", while `2` means the program never got far enough to give an answer.

### TR-3: Anagram check is order-independent

The anagram check must give the same answer regardless of the order of its two arguments. `<run> anagram "a" "b"` and `<run> anagram "b" "a"` must agree. A sort-based approach (normalize, sort, compare) satisfies this naturally; a frequency-count approach does too.

### TR-4: Unicode handling is defined, not assumed

The spec only requires correctness on ASCII letters. The program must not crash on non-ASCII input; it may either treat non-ASCII characters as non-letters (drop them during normalization) or extend case folding to them — either is acceptable as long as the choice is consistent across both checks.

## Acceptance Criteria

- AC-1: `<run> palindrome "A man a plan a canal Panama"` prints a match verdict and exits with code `0`.
- AC-2: `<run> palindrome ""` (empty string) prints a match verdict and exits with code `0`.
- AC-3: `<run> palindrome "hello"` prints a no-match verdict and exits with code `1`.
- AC-4: `<run> anagram "listen" "Silent"` prints a match verdict and exits with code `0`.
- AC-5: `<run> anagram "conversation" "voices rant on"` prints a match verdict and exits with code `0`.
- AC-6: `<run> anagram "hello" "world"` prints a no-match verdict and exits with code `1`.
- AC-7: `<run> anagram "a" "b"` and `<run> anagram "b" "a"` produce the same exit code.
- AC-8: `<run>` with no arguments, an unknown subcommand, or the wrong number of arguments prints a usage message and exits with code `2`.
