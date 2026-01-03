# AGENTS.md

Guide for AI agents working on the mutO codebase.

## Quick start

```bash
# Run tests (GOCACHE must be absolute path)
GOCACHE=$PWD/tmp/go-cache go test ./...

# Build CLI
make build

# Run a file
./build/muto run tmp/main.mu

# Run with mutation steps shown
./build/muto run --explain tmp/main.mu

# Start REPL
./build/muto repl
```

## Agent workflow

After any code edit, run tests unless the user says otherwise:
```bash
GOCACHE=$PWD/tmp/go-cache go test ./...
```

## Core concepts

### Objects and nodes
- **Node**: Any value - can be a class, number, string, boolean, tag, or object
- **Object**: A head (node) plus parameters. Written as `head param1 param2`
- **Compound object**: Nested application like `((f 1) 2)` stored internally as head `f` + param chain `[[1], [2]]`

### Patterns (for rule matching)
- **Determinant**: The part that determines which rule matches - the class name at the head
- **Non-determinant**: Patterns in parameters that don't affect rule selection
- **Param chain**: The flattened list of parameter groups from nested objects
- **Conjunction (`^`)**: Binds intermediate objects during matching. `(f X)^P Y` matches `(f 1) 2` with `P=(f 1)`

### Rules and mutation
- **Normal rule**: `name pattern = result` - matches after children are mutated
- **Active rule**: `@ name pattern = result` - matches before children mutate (higher priority)
- **Mutation**: One rewrite step. Program runs until no more mutations apply.

### Portal
- IO operations use "portal" - a registry of named ports (`stdout`, `stdin`, `spawner`)
- `print!`, `input!`, `spawn!` are sugar for portal calls

## Key entry points

| Area | Entry point |
|------|-------------|
| CLI | `cmd/cli/cli/main.go` |
| Program build | `builder/program/main.go` |
| Mutation loop | `program/main.go` |
| Rule building | `core/mutation/rule/builder/main.go` |
| Pattern extraction | `core/mutation/rule/extractor/main.go` |
| Parsing rules | `parser/file/rule.go` |
| Parsing patterns | `parser/pattern/determinant.go` |
| Builtins (code) | `builtin/global/code.go` |
| Builtins (Go) | `builtin/global/mutator.go` |
| Portal devices | `builtin/portal/default.go` |

## Common tasks

### Add a builtin rule (muto code)
1. Create file in `builtin/global/` with the rule string
2. Register in `builtin/global/code.go`

### Add a builtin rule (Go mutator)
1. Implement `mutator.NamedUnit` interface
2. Register in `builtin/global/mutator.go`

### Add a portal device
1. Implement `portal.Port` interface (see `builtin/portal/spawner.go`)
2. Register in `builtin/portal/default.go`

### Add a builtin module
1. Create module in `builtin/<name>/module.go` (see `builtin/time/`)
2. Map in `builtin/module.go`

### Add pattern syntax
1. Parser: `parser/pattern/determinant.go` or `parser/pattern/param_part.go`
2. AST node: `syntaxtree/pattern/`
3. Extractor: `core/mutation/rule/extractor/` or `core/pattern/extractor/`

### Add a CLI command
1. Create in `cmd/cli/<name>/cli.go` (see `cmd/cli/run/`)
2. Wire in `cmd/cli/cli/main.go`

### Add a REPL command
1. Create command in `builder/repl/core/command/`
2. Handle in `builder/repl/core/executor/main.go`

## Directory overview

```
cmd/cli/          CLI entry point and commands
builder/          Builds programs from syntax trees
  program/        File execution builder
  repl/           REPL builder and commands
builtin/          Built-in rules and modules
  global/         Global builtins (code.go, mutator.go, *_ops.go)
  portal/         IO devices (stdout, stdin, spawner)
  time/           Time module
core/             Runtime core
  base/           Nodes, objects, mutation primitives
  module/         Module system
  mutation/rule/  Rule building and extraction
  pattern/        Pattern extractors
  portal/         Portal interface
parser/           Parsing
  file/           Module and rule parsing
  pattern/        Pattern parsing
  result/         Result expression parsing
syntaxtree/       AST nodes
  pattern/        Pattern AST (conjunction, nesting)
  result/         Result AST
program/          Mutation loop
examples/         Example programs
```

## Gotchas

- `NewCompoundObject` panics if the head is itself an object
- User-defined modules support only one file (panics otherwise)
- Portal `.call` requires exactly one param and string key
- `spawn!` requires at least one param; `(spawn!)` alone won't mutate
- `sleep` expects a number argument
- `try` returns `.empty` tag when child mutation fails
- Structure `.get`/`.set` only works with tag heads
- Active rules always have priority over normal rules
- Mutation loop halts when `Mutate()` returns empty

## Test locations

Key test files:
- `builder/program/main_test.go` - Integration tests
- `core/mutation/rule/extractor/main_test.go` - Pattern extraction
- `parser/pattern/determinant_test.go` - Pattern parsing
- `builtin/global/*_test.go` - Builtin operations
