# Number Guessing Game

An interactive command-line game, one round per run: the program picks a secret integer, the user guesses until correct, and the program reports how many guesses it took.

## Functional Requirements

### FR-1: Pick a secret in range

On startup the system picks a secret integer uniformly at random from `[lower, upper]` — every value including both endpoints is equally likely — and never reveals it. Defaults are `1` and `100`; `--lower N` and `--upper N` override either bound. The active range is printed at startup. A range where `lower > upper` is a usage error: it is reported before any prompt and exits with code `2`.

### FR-2: Guess until correct

The system reads one guess per line from standard input and answers each with exactly one of three outcomes: the secret is *higher*, the secret is *lower*, or the guess is *correct*. On a correct guess it prints how many valid guesses the round took and exits with code `0`.

### FR-3: Invalid guesses are rejected, not fatal

A line that isn't an integer, or an integer outside the range, gets a short error message and a reprompt. It does not terminate the program and does not count toward the guess total.

### FR-4: Clean end without a win

If input ends (end-of-file) before a correct guess, the program exits cleanly — no crash output or stack trace — with exit code `1`.

## Acceptance Criteria

- AC-1: `<run>` with no arguments prints the active range, accepts integer guesses, and reports a guess count on win.
- AC-2: `<run> --lower 1 --upper 1000` draws a secret in `[1, 1000]` and accepts guesses up to 1000.
- AC-3: Blank lines, letters, or symbols at the prompt produce a reprompt rather than terminating the program.
- AC-4: A round with several invalid lines that ends on the *N*th valid guess reports exactly *N* as the guess count.
- AC-5: With `--lower 1 --upper 10`, entering `42` triggers a reprompt and does not count toward the guess total.
- AC-6: `<run> --lower 50 --upper 10` prints a usage error and exits with code `2` without prompting for any guesses.
- AC-7: `<run> < /dev/null` prints no panic or stack trace and exits with a non-zero status.
