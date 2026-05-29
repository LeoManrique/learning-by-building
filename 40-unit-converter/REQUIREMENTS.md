# Unit Converter

A one-shot command-line tool that converts a single number between two units of the same family — temperature (Celsius, Fahrenheit, Kelvin) plus one more family (length or mass). The user is anyone who needs a quick conversion at the terminal. The tool reads its arguments, runs a formula, prints the result, and exits. No state, no interactive loop, no files, no network.

## Functional Requirements

### FR-1: Convert a value between two units of the same family
Given a number and two unit labels from the same family, the tool prints the converted value on a single line to stdout.

### FR-2: Support at least two unit families
The tool supports temperature (Celsius, Fahrenheit, Kelvin) and at least one more family (length or mass). Each family has at least three units identified by short labels (e.g. `C`, `F`, `K`, or `m`, `km`, `mi`).

### FR-3: Conversion stays inside one family
The source and target units must be from the same family. Converting across families (e.g. Celsius to meters) is rejected as a user error.

### FR-4: Pick the number of decimals
A flag sets how many decimal places appear in the printed result. The default is 2.

### FR-5: Reject unknown units with a usage message
If a unit label isn't supported, the tool prints a usage message naming the bad input and exits with a non-zero code. No number is printed.

## Technical Requirements

### TR-1: One-shot process
Each run parses its arguments, does one conversion, prints one line, and exits. No interactive loop, no files, no shared state between runs.

### TR-2: Use 64-bit floats for the math
Numbers are stored internally as 64-bit floats. The decimal-places setting only affects the printed output — the math itself uses full float precision so rounding doesn't pile up.

### TR-3: Only the listed units are accepted
The set of supported units is fixed and written out in the code. Any label not on the list is rejected as a user error — never guessed at or defaulted. Why: keeping the list closed means typos fail loudly at parse time instead of producing a silently-wrong number.

### TR-4: Exit codes tell success from user error
Exit 0 means the conversion worked. Exit 2 means the user got something wrong (unknown unit, mixed families, missing argument, value that isn't a number, bad flag). Why: a distinct non-zero code lets shells and scripts react to bad input without reading stderr.

### TR-5: Round-trip stays stable within the chosen precision
For any supported pair of units A and B and any reasonable value v, converting v from A to B and then back from B to A returns v within the chosen precision. This has to hold for the formulas themselves, not just the printer.

### TR-6: Numbers on stdout, errors on stderr
Successful output goes to stdout. Usage messages and error explanations go to stderr. Why: a caller piping the result into another tool sees only the number on stdout.

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
