# Number Guessing Game — Requirements

A small interactive command-line game. The program picks a secret integer, then repeatedly reads a guess from the user and reports whether the secret is *higher*, *lower*, or *equal* to the guess. The round ends when the user guesses correctly, at which point the program reports how many guesses it took and exits.

The application is invoked from a shell as a single, short-lived process. There is no persistent state, no configuration file, and no networking — one run of the program corresponds to exactly one round of the game.

## Functional Requirements

### FR-1: Start a round with a secret in range
On startup, the system selects a secret integer uniformly at random from a closed integer interval `[MIN, MAX]`. The chosen secret is held in memory for the duration of the round and is never echoed to the user.

### FR-2: Prompt and read guesses
The system prints a prompt and reads one line of input from standard input per guess. Each line is interpreted as a single integer guess. The system continues prompting until the user wins or input is exhausted.

### FR-3: Respond to each guess
For every successfully parsed guess, the system prints exactly one of three outcomes: the secret is *higher* than the guess, *lower* than the guess, or the guess is *correct*. The wording is consistent across guesses so the user can trust the signal.

### FR-4: Tolerate invalid input without crashing
If the input line is blank, contains non-numeric characters, or is otherwise unparseable as an integer, the system prints a short error message and reprompts. Invalid input does not terminate the program and does not count as a guess. *Why:* the user is interacting in real time and a single typo should not abort their round.

### FR-5: Report guess count on win
When the user guesses correctly, the system prints the total number of valid guesses it took to reach the secret, then exits successfully. Only successfully parsed numeric guesses are counted; invalid input is excluded from the total.

### FR-6: Configurable range via flag
The user may supply a flag (for example `--range MIN MAX`) to widen or otherwise change the secret range before the round starts. When the flag is omitted, the system uses a built-in default range (e.g. `1..100`). The chosen range is communicated to the user at startup so they know the bounds they're guessing within.

## Technical Requirements

### TR-1: Range validation
The system validates the supplied range before selecting the secret. `MIN` and `MAX` must both parse as integers and `MIN` must not exceed `MAX`. A range with `MIN == MAX` is permitted (the round is trivial but well-defined). Any violation is reported as a usage error before any prompt is shown.

### TR-2: Uniform secret selection
The secret is drawn uniformly from the closed interval `[MIN, MAX]` — every integer in the range, including both endpoints, has equal probability of being chosen. *Why:* off-by-one mistakes here (excluding an endpoint, or skewing the distribution) silently bias the game; the requirement is the contract that prevents that.

### TR-3: Line-buffered input loop
Input is consumed one line at a time. The system does not require any specific line ending and treats the trailing newline as terminator only, not as part of the guess. Leading and trailing whitespace on the line is ignored before parsing.

### TR-4: Out-of-range guesses
A successfully parsed guess that falls outside `[MIN, MAX]` is treated as invalid input under FR-4: the system reprompts and does not count it as a guess. *Why:* such guesses cannot possibly be correct, and counting them would let the user inflate or game the guess total accidentally.

### TR-5: End-of-input handling
If standard input reaches end-of-file before the user wins (e.g. the input stream is closed or the user sends EOF), the system terminates cleanly without a stack trace or panic. It exits with a non-zero status to signal that the round did not end in a win.

### TR-6: Exit codes
Exit codes are meaningful so the program composes cleanly in shell pipelines:
- `0` — the user guessed correctly and the round ended normally
- `1` — the round ended without a win (EOF, user abort)
- `2` — usage error (bad flag, malformed range, `MIN > MAX`)

### TR-7: No persistence, no shared state
The program holds no state across invocations. Each run is a fresh round with a freshly drawn secret and a fresh guess counter. There is no save file, no history, and no cross-round statistics.

## Acceptance Criteria

The application is considered complete when all of the following hold:

- **AC-1 — Default range round:** Running `<run>` with no arguments prints the active range, accepts integer guesses, and ends with a `correct` message and a guess count once the user enters the secret.
- **AC-2 — Range flag widens the range:** Running `<run> --range 1 1000` (or the chosen flag form) draws a secret somewhere in `[1, 1000]`, and guesses outside `[1, 100]` are not rejected as out-of-range.
- **AC-3 — Invalid input survival:** Entering blank lines, letters, or symbols at the prompt produces a reprompt rather than terminating the program. After several invalid lines the user can still enter a valid number and proceed.
- **AC-4 — Invalid input is not counted:** A round in which the user enters several invalid lines and then wins on the *N*th valid numeric guess reports exactly *N* as the guess count.
- **AC-5 — Out-of-range guess is rejected:** With a range of `[1, 10]`, entering `42` triggers a reprompt and does not count toward the guess total.
- **AC-6 — Bad range exits with usage code:** Running `<run> --range 50 10` (`MIN > MAX`) prints a usage error and exits with code `2` without prompting for any guesses.
- **AC-7 — EOF exits cleanly:** Piping an empty input (e.g. `<run> < /dev/null`) prints no panic or stack trace and exits with a non-zero status.
- **AC-8 — Distribution sanity:** Running the program many times with a small range (e.g. `[1, 5]`) and recording the secrets shows every value in the range appearing — no value is systematically excluded.
