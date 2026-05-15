# Number Guessing Game — Requirements

A small interactive command-line game. The program picks a secret integer, then repeatedly reads guesses from the user and reports whether the secret is *higher*, *lower*, or *equal* to the guess. The round ends when the user guesses correctly, at which point the program reports how many guesses it took and exits.

The application runs as a single short-lived process: one invocation, one round, no persistence.

## Functional Requirements

### FR-1: Start a round with a secret in range
On startup, the system selects a secret integer uniformly at random from a closed integer interval `[lower, upper]`. The secret is held in memory for the round and is never echoed to the user.

### FR-2: Read guesses interactively
The system reads one line of input from standard input per guess and interprets it as an integer. It continues reading until the user wins or input is exhausted.

### FR-3: Respond to each guess
For every successfully parsed guess the system prints exactly one of three outcomes: the secret is *higher* than the guess, *lower* than the guess, or the guess is *correct*.

### FR-4: Tolerate invalid input without crashing
If the input line isn't a valid integer (blank, non-numeric, etc.), the system prints a short error message and reprompts. Invalid input does not terminate the program and does not count as a guess.

### FR-5: Report guess count on win
When the user guesses correctly, the system prints the total number of valid guesses it took to reach the secret and exits successfully.

### FR-6: Configurable range via flags
The user may supply `--lower N` and/or `--upper N` flags to set the bounds. Either flag may be omitted, in which case a built-in default is used (e.g. `lower = 1`, `upper = 100`). The chosen range is communicated to the user at startup so they know the bounds they're guessing within.

## Technical Requirements

### TR-1: Range validation
The system validates the supplied range before drawing the secret. `lower` must not exceed `upper`; equal values are permitted (trivial round). A bad range is reported as a usage error before any prompt is shown.

### TR-2: Uniform secret selection
The secret is drawn uniformly from the closed interval `[lower, upper]` — every integer in the range, *including both endpoints*, has equal probability. *Why:* off-by-one mistakes here (excluding an endpoint) silently bias the game without surfacing as a visible bug.

### TR-3: Out-of-range guesses are invalid
A successfully parsed guess outside `[lower, upper]` is treated as invalid input under FR-4: the system reprompts and does not count it as a guess.

### TR-4: End-of-input handling
If standard input reaches end-of-file before the user wins, the system terminates cleanly without a stack trace or panic and exits with a non-zero status to signal that the round did not end in a win.

### TR-5: Exit codes
- `0` — the user guessed correctly
- `1` — round did not end in a win (EOF, abort)
- `2` — usage error (bad flag, `lower > upper`)

## Acceptance Criteria

- **AC-1 — Default round:** `<run>` with no arguments prints the active range, accepts integer guesses, and reports a guess count on win.
- **AC-2 — Range flags widen the range:** `<run> --lower 1 --upper 1000` draws a secret in `[1, 1000]` and accepts guesses up to 1000.
- **AC-3 — Invalid input survival:** Blank lines, letters, or symbols at the prompt produce a reprompt rather than terminating the program.
- **AC-4 — Invalid input is not counted:** A round in which the user enters several invalid lines and then wins on the *N*th valid numeric guess reports exactly *N* as the guess count.
- **AC-5 — Out-of-range guess is rejected:** With `--lower 1 --upper 10`, entering `42` triggers a reprompt and does not count toward the guess total.
- **AC-6 — Bad range exits with usage code:** `<run> --lower 50 --upper 10` prints a usage error and exits with code `2` without prompting for any guesses.
- **AC-7 — EOF exits cleanly:** `<run> < /dev/null` prints no panic or stack trace and exits with a non-zero status.
