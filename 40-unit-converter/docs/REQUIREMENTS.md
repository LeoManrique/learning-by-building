# Unit Converter

A one-shot command-line tool: read the arguments, convert one number between two units of the same family, print the result, exit. No state, no files, no network.

## Functional Requirements

### FR-1: Convert within a family

`<run> --from <unit> --to <unit> <value>` prints the converted value on a single line to stdout and exits with code `0`. Two families are supported — temperature (`C`, `F`, `K`) and one more (length or mass) — each with at least three units identified by short labels.

### FR-2: Only listed units, only same-family pairs

Only the listed unit labels are accepted, and both units must belong to the same family. An unknown label or a cross-family pair (e.g. Celsius to meters) is a user error — like any other user error (missing argument, non-numeric value, bad flag) it prints a message naming the bad input on stderr, prints no number on stdout, and exits with code `2`.

### FR-3: Choose the printed precision

`--precision N` sets how many decimal places are printed; the default is 2. It affects only the printing — the math runs at full precision, so converting a value from A to B and back from B to A returns the original value within the displayed precision.

## Acceptance Criteria

- AC-1: `<run> --from C --to F 100` exits 0 and prints `212.00`.
- AC-2: `<run> --from C --to K 100` exits 0 and prints `373.15`.
- AC-3: `<run> --from km --to m 1` exits 0 and prints `1000.00` (or the matching result for the chosen second family).
- AC-4: `<run> --from C --to km 100` exits with code 2 and prints a message saying the units are from different families. No number on stdout.
- AC-5: `<run> --from X --to F 100` exits with code 2 and prints a message naming the unknown unit `X`.
- AC-6: `<run> --precision 4 --from C --to F 100` exits 0 and prints `212.0000`.
- AC-7: `<run> --from C --to F 0` exits 0 and prints `32.00`.
- AC-8: For any supported pair, converting `v` from A to B and then back from B to A returns `v` within the displayed precision (e.g. `<run> --from C --to F 100` → `212.00`, then `<run> --from F --to C 212.00` → `100.00`).
- AC-9: `<run>` with no arguments, or with a non-number where a number is expected, exits with code 2 and prints a usage message on stderr.
