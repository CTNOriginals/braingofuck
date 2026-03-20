# BrainGoFuck

A [Brainfuck](https://esolangs.org/wiki/Brainfuck) interpreter in Go — **stage 1**. Source is tokenized with line and column positions, then executed on a byte tape with a stack for loop jumps. The `compiler` package name reflects how the project is organized; execution is interpretation from tokens, not compilation to native code.

## Features

- **Tokenizer** (`tokenizer/`) — maps `><+-.,[]` to typed tokens; records `Line` and `Col` per token; other characters are ignored (comments and formatting are fine).
- **Tape** (`compiler/`) — fixed number of **byte** cells (`Cell`); pointer **wraps** at both ends (circular tape).
- **Loops** — `[` always pushes the token index; `]` jumps back while the current cell is non-zero, otherwise pops and continues (see command table for the non-standard `[` behavior).
- **Output** — `.` appends the current cell as a byte; accumulated output is printed when the run finishes.
- **Samples** — see `brainfuck/hello.bf`, `brainfuck/loops.bf`, and `brainfuck/proto.bf`.

## Requirements

- [Go](https://go.dev/dl/) **1.25** or compatible toolchain (see `go.mod`).
- Module `github.com/CTNOriginals/CTNGoUtils/v2` (fetched with `go mod download` / first build).

## Quick start

From the repository root:

```bash
go run .
```

Or:

```bash
make run
```

**Note:** The Brainfuck source path is **hardcoded** in [`main.go`](main.go) (`./brainfuck/hello.bf`, with `./brainfuck/proto.bf` commented as an alternative). To run another file, change that path and rebuild/run.

## Project layout

```
main.go              # entry: load file, tokenize, execute
tokenizer/           # lexer → token stream with positions
compiler/            # tape, cells, loop stack, execution
brainfuck/           # sample .bf programs
Makefile             # convenience targets
go.mod / go.sum      # module and dependencies
LICENSE              # MIT
```

## Brainfuck commands

| Char | Meaning |
|------|---------|
| `>` | Move pointer right (wraps at end of tape) |
| `<` | Move pointer left (wraps at start) |
| `+` | Increment current cell (byte wrap) |
| `-` | Decrement current cell (byte wrap) |
| `.` | Output current cell as a byte |
| `,` | Input (see limitations) |
| `[` | Push the current token index onto the loop stack (always; no “skip if zero” at `[`) |
| `]` | If current cell is non-zero, jump back to the `[` index on the stack top; if zero, pop the stack and continue after `]` |

Note: Classic Brainfuck skips the body when the cell is zero at `[`. This stage **always** enters the code after `[` until a `]` decides whether to repeat; many programs still behave as expected, but idioms that rely on skipping a zero-depth `[` may differ.

## Makefile

| Target | Description |
|--------|-------------|
| `make run` | `go run .` |
| `make wrun` | Watch `.go` and `.bf` with [wgo](https://github.com/bokwoon95/wgo) and re-run (install `wgo` separately) |

The `asci` target is a local maintainer helper (hardcoded path to another tool on disk); it is not part of the public workflow for this project.

## Stage 1 limitations and next steps

- **`,` (input)** — `INP` in `compiler/compiler.go` is not implemented yet (no-op / TODO).
- **Configuration** — tape size is passed from `main` as `256` cells; there is no CLI for path or size yet.
- **Debug output** — execution prints verbose traces (e.g. state when a token is at column 1, loop jump logging). A later stage may gate this behind a flag.
- **Tests** — no automated tests in this stage.

## License

[MIT](LICENSE). Copyright (c) 2026 CTNOriginals.
