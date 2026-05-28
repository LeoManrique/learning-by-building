# Unit Converter

A one-shot command-line tool that converts a single numeric value between two units of the same family — temperature (Celsius, Fahrenheit, Kelvin) plus one additional family (length or mass). The user is anyone who needs a quick conversion at the terminal. The system is a short-lived CLI process: it parses argv, applies a formula, prints the result, and exits. There is no state, no interactive loop, no persistence, and no network.

## Functional Requirements

### FR-1: Convert a value between two units of the same family
Given a numeric value and two unit labels belonging to the same family, the system computes the converted value and prints it to stdout on a single line.

### FR-2: Support at least two unit families
The system supports temperature (Celsius, Fahrenheit, Kelvin) and at least one additional family (length or mass). Each family contains at least three units identified by short labels (e.g. `C`, `F`, `K`, or `m`, `km`, `mi`).

### FR-3: Conversion only within a single family
Source and target units must belong to the same family. A request to convert across families (e.g. Celsius to meters) is rejected as a user error.

### FR-4: Accept both positional and flag invocation shapes
The CLI accepts either positional form (`<run> 100 C F`) or named-flag form (`<run> --from C --to F 100`). Both shapes produce identical output for identical inputs.

### FR-5: Configurable output precision
A flag controls the number of decimal places in the printed result. The default is 2 decimals.

### FR-6: Reject unknown units with a usage message
If either unit label does not match any supported unit, the system prints a usage message naming the offending input and exits with a non-zero code without producing a numeric result.

## Technical Requirements

### TR-1: Stateless one-shot process
Each invocation parses argv, performs a single conversion, prints one line, and exits. There is no interactive loop, no file I/O, and no state shared across invocations.

### TR-2: 64-bit float internal representation
Values are represented internally as 64-bit floating-point numbers. The configured precision applies only at the formatting step — the conversion arithmetic itself uses full float precision so that intermediate rounding does not accumulate.

### TR-3: Exhaustive matching over a closed unit set
Unit label resolution matches inputs against a closed, enumerated set of known units. Every supported unit is named explicitly; an unmatched label is treated as a user error, never silently coerced or defaulted. Why: the closed match makes the supported set readable directly from the code and catches typos at parse time rather than producing a silently-wrong number.

### TR-4: Exit codes distinguish success from user error
Exit code 0 indicates a successful conversion. Exit code 2 indicates a user error (unknown unit, cross-family request, missing argument, unparseable numeric value, malformed flag, or unknown flag). Why: a distinct non-zero code lets shells and scripts react to bad input without parsing stderr.

### TR-5: Round-trip stability within the configured precision
For any supported pair of units A and B and any reasonable value v, converting v from A to B and then converting that result back from B to A yields the original value v within the configured precision. This is a property of the formulas, not just the printer — the internal arithmetic must preserve the value across an inverse pair.

### TR-6: All output on stdout; errors on stderr
Successful numeric output goes to stdout. Usage messages and error explanations go to stderr. Why: callers piping the result into another tool see only the number on stdout.

## Acceptance Criteria

- AC-1: `<run> 100 C F` exits 0 and prints `212.00`.
- AC-2: `<run> --from C --to F 100` exits 0 and prints `212.00` (alternate invocation shape, same result).
- AC-3: `<run> 100 C K` exits 0 and prints `373.15`.
- AC-4: `<run> 1 km m` exits 0 and prints `1000.00` (or the analogous result for the chosen second family).
- AC-5: `<run> 100 C km` exits with code 2 and prints a message indicating the units belong to different families. No numeric result is printed on stdout.
- AC-6: `<run> 100 X F` exits with code 2 and prints a message naming the unknown unit `X`.
- AC-7: `<run> --precision 4 100 C F` exits 0 and prints `212.0000`.
- AC-8: `<run> 0 C F` exits 0 and prints `32.00`.
- AC-9: For any supported pair, the value obtained by converting `v` from A to B and then back from B to A equals `v` within the displayed precision (e.g. `<run> 100 C F` → `212.00`, then `<run> 212.00 F C` → `100.00`).
- AC-10: `<run>` with no arguments, or with a non-numeric value where a number is expected, exits with code 2 and prints a usage message on stderr.
