## Project snapshot
- mutO interpreter/CLI with run + repl commands wired in `main.go`, `cmd/cli/cli/main.go`, `cmd/cli/run/cli.go`, and `cmd/cli/repl/cli.go`.
- Program builder assembles global/builtin modules, portal, and user module in `builder/program/main.go`, `builtin/global/module.go`, `builtin/portal/default.go`, and `core/module/user_defined.go`.
- Mutation loop and initial object live in `program/main.go`.
- Builtins split into global rules/mutators and portal devices in `builtin/global/code.go`, `builtin/global/mutator.go`, and `builtin/portal/default.go`.

## Quick start (tested)
- Tests (worked here): `GOCACHE=tmp/go-cache go test ./...` using module config in `go.mod` and cache dir `tmp/go-cache`.
- Build CLI: `make build` from `Makefile` outputs `build/muto` per `Makefile`.
- Run sample file: `make run` uses `tmp/main.mu` via `Makefile`.
- Explain mode: `make run-explain` prints mutation steps via `cmd/cli/run/execute.go` and `program/main.go`.
- REPL: `make repl` runs CLI REPL wired in `cmd/cli/repl/cli.go`.

## Makefile targets
- `make build` -> `go build -o build/muto ./cmd/cli` in `Makefile`.
- `make build-cli` -> `go build -o build/muto ./cmd/cli` in `Makefile`.
- `make run` -> `muto run tmp/main.mu` in `Makefile`.
- `make run-explain` -> `muto run --explain tmp/main.mu` in `Makefile`.
- `make run-tmp-main` -> `muto run tmp/main.mu` in `Makefile`.
- `make run-tmp-main-explain` -> `muto run --explain tmp/main.mu` in `Makefile`.
- `make repl` -> `muto repl` in `Makefile`.
- `make repl-tmp-main` -> `muto repl` in `Makefile`.
- `make example-tictactoe` -> `muto run examples/tictactoe.mu` in `Makefile`.
- `make go-get-common` uses curl/jq for dependency update in `Makefile`.

## CLI commands
- CLI app registers `run` and `repl` commands in `cmd/cli/cli/main.go`.
- `run` requires a filename and prints an error if missing in `cmd/cli/run/cli.go`.
- `run --explain` toggles the after-mutation hook in `cmd/cli/run/execute.go`.
- File execution reads bytes then builds program in `cmd/cli/run/execute.go` and `builder/program/main.go`.
- REPL wiring constructs reader/printer around readline in `cmd/cli/repl/cli.go`.
- REPL loop exits on quit command return code in `cmd/cli/repl/repl.go` and `builder/repl/core/executor/main.go`.
- REPL parsing uses `parser/repl` and maps to commands in `builder/repl/core/command/parser.go` and `parser/repl/statement.go`.
- CLI entrypoint function is `main.go`.

## Repo map
- `build/` is the CLI output directory referenced by `Makefile` and `shellstartup.sh`.
- `builder/program` builds programs from syntax trees in `builder/program/main.go`.
- `builder/repl` bootstraps REPL program in `builder/repl/main.go`.
- `builder/repl/core/command` defines REPL commands in `builder/repl/core/command/mutate_node.go`.
- `builder/repl/core/executor` dispatches commands in `builder/repl/core/executor/main.go`.
- `builder/repl/core/program` wraps `program.Program` in `builder/repl/core/program/main.go`.
- `builder/repl/core/reader` handles input/errors in `builder/repl/core/reader/reader.go`.
- `builtin/global` hosts builtin rule strings in `builtin/global/code.go`.
- `builtin/global` Go mutators live in `builtin/global/mutator.go` and `builtin/global/number_ops.go`.
- `builtin/portal` defines portal devices in `builtin/portal/stdout.go` and `builtin/portal/stdin.go`.
- `builtin/portal` registers defaults in `builtin/portal/default.go`.
- `builtin/time` module and sleep mutator are in `builtin/time/module.go` and `builtin/time/sleep.go`.
- builtin import mapping is in `builtin/module.go`.
- `cmd/cli` app setup is in `cmd/cli/cli/main.go`.
- `cmd/cli/run` command and execute path are in `cmd/cli/run/cli.go` and `cmd/cli/run/execute.go`.
- `cmd/cli/repl` command and loop are in `cmd/cli/repl/cli.go` and `cmd/cli/repl/repl.go`.
- `common/fn` helpers are in `common/fn/compose.go`.
- `common/slc` helpers are in `common/slc/map.go`.
- `common/strutil` helpers are in `common/strutil/stringer.go`.
- `common/typ` helpers are in `common/typ/zero.go`.
- `core/base` defines nodes/objects/classes in `core/base/node.go` and `core/base/compound_object.go`.
- `core/module` defines modules/dependencies in `core/module/base.go` and `core/module/dependency.go`.
- `core/mutation/rule` and builder are in `core/mutation/rule/main.go` and `core/mutation/rule/builder/main.go`.
- `core/mutation/rule/mutator` collection/switches are in `core/mutation/rule/mutator/collection.go` and `core/mutation/rule/mutator/switch.go`.
- `core/portal` interfaces and store are in `core/portal/port.go` and `core/portal/main.go`.
- `parser/file` module/rule parsing is in `parser/file/module.go` and `parser/file/rule.go`.
- `parser/pattern` pattern parsing is in `parser/pattern/determinant.go` and `parser/pattern/param_part.go`.
- `parser/result` object/structure parsing is in `parser/result/object.go` and `parser/result/structure.go`.
- `parser/repl` statement parsing is in `parser/repl/statement.go`.
- `program` mutation loop and initial object are in `program/main.go`.
- `syntaxtree` AST nodes are in `syntaxtree/module.go`, `syntaxtree/rule.go`, and `syntaxtree/active_rule.go`.
- `syntaxtree/result` nodes are in `syntaxtree/result/object.go`.
- `examples` sample program lives in `examples/tictactoe.mu`.
- `tmp` task artifacts include `tmp/plan.md` and cache `tmp/go-cache`.
- `README.md` documents language concepts and CLI usage in `README.md`.

## Execution flows
- Run flow: `main.go` -> `cmd/cli/cli/main.go` -> `cmd/cli/run/cli.go` -> `cmd/cli/run/execute.go`.
- Program build flow: `cmd/cli/run/execute.go` -> `builder/program/main.go` -> `builtin/global/module.go` -> `builtin/portal/default.go` -> `core/module/user_defined.go`.
- Program mutation loop: `program/main.go` -> `core/base/mutatable_node.go` -> `core/base/rule_based_class.go`.
- Portal call flow: `builtin/global/portal.go` -> `core/portal/main.go` -> `builtin/portal/stdout.go`.
- Print flow: `builtin/global/stdio.go` -> `builtin/global/portal.go` -> `builtin/portal/stdout.go`.
- Input flow: `builtin/global/stdio.go` -> `builtin/global/portal.go` -> `builtin/portal/stdin.go`.
- Spawn flow: `builtin/global/spawn.go` -> `builtin/global/portal.go` -> `builtin/portal/spawner.go` -> `core/base/mutatable_node.go`.
- REPL flow: `cmd/cli/repl/cli.go` -> `builder/repl/main.go` -> `builder/repl/core/reader/reader.go` -> `builder/repl/core/executor/main.go`.
- REPL mutate node: `builder/repl/core/command/mutate_node.go` -> `builder/repl/core/program/main.go` -> `program/main.go`.
- REPL add rule: `builder/repl/core/command/add_rule.go` -> `builder/repl/core/program/main.go` -> `core/module/base.go`.

## Portal and IO map
- `print!` uses `portal (.call "stdout" X)` in `builtin/global/stdio.go` and port mapping in `builtin/portal/default.go`.
- `input!` uses `portal (.call "stdin" $)` in `builtin/global/stdio.go` and `builtin/portal/default.go`.
- `spawn!` uses `portal (.call "spawner" (X Xs...))` in `builtin/global/spawn.go` and `builtin/portal/default.go`.
- Portal `.call` requires string key and value param in `builtin/global/portal.go`.
- Portal interface is `Port.Call` in `core/portal/port.go`.
- Portal storage and lookup is in `core/portal/main.go`.
- Stdout port prints string nodes in `builtin/portal/stdout.go`.
- Stdin port reads when passed `$` class in `builtin/portal/stdin.go` and `core/base/convention.go`.
- Spawner port runs mutable nodes in a goroutine in `builtin/portal/spawner.go` and `core/base/mutatable_node.go`.

## Mutation and rules
- Rule-based class mutation tries active then normal in `core/base/rule_based_class.go`.
- Active rules are parsed with `@` in `parser/file/rule.go` and represented by `syntaxtree/active_rule.go`.
- Normal rule build path is `core/mutation/rule/main.go` -> `core/mutation/rule/builder/main.go`.
- Pattern matching uses extractors in `core/mutation/rule/extractor/main.go` and `core/mutation/rule/extractor/top_level.go`.
- Param chain mutation walks children in `core/base/children.go`.
- Param chain behavior (append/chain/slice) is in `core/base/param_chain.go`.
- Mutation result handling uses `core/base/mutation.go`.
- Conventional `$` class and tags are in `core/base/convention.go`.
- Structures support `.get`/`.set` tags in `core/base/structure.go` and `core/base/convention.go`.
- `try` mutator returns `.empty` or `.value` tags in `builtin/global/object_ops.go` and `core/base/convention.go`.
- Global builtin rules from code strings are loaded in `builtin/global/code.go`.
- `spawn!` returns `$` via `base.Null()` in `builtin/portal/spawner.go` and `core/base/convention.go`.

## Parsing and AST
- Module parsing from string is in `parser/file/module.go`.
- Rule parsing is in `parser/file/rule.go`.
- Pattern determinant parsing is in `parser/pattern/determinant.go`.
- Pattern param part parsing is in `parser/pattern/param_part.go`.
- Object result parsing is in `parser/result/object.go`.
- Structure result parsing is in `parser/result/structure.go`.
- REPL statement parsing is in `parser/repl/statement.go`.
- Module AST nodes are in `syntaxtree/module.go`.
- Rule AST nodes are in `syntaxtree/rule.go` and `syntaxtree/active_rule.go`.
- Result AST nodes are in `syntaxtree/result/object.go`.

## REPL details
- REPL builder attaches global module and portal in `builder/repl/main.go`.
- REPL reader reads lines and prints errors in `builder/repl/core/reader/reader.go`.
- REPL parser maps statements to commands in `builder/repl/core/command/parser.go`.
- REPL mutate-node command struct is in `builder/repl/core/command/mutate_node.go`.
- REPL add-rule command struct is in `builder/repl/core/command/add_rule.go`.
- REPL quit command struct is in `builder/repl/core/command/quit.go`.
- REPL executor dispatch is in `builder/repl/core/executor/main.go`.
- REPL program wrapper uses `program.Program` in `builder/repl/core/program/main.go` and `program/main.go`.
- REPL command types are in `builder/repl/core/command/abstract.go`.
- REPL command parsing uses repl AST nodes in `builder/repl/core/command/parser.go` and `syntaxtree/repl/statement.go`.

## Extension points
- Add a global rule by creating a code string file like `builtin/global/spawn.go` and registering it in `builtin/global/code.go`.
- Add a Go mutator implementing `mutator.NamedUnit` and register in `builtin/global/mutator.go`.
- Add a portal device by implementing `portal.Port` like `builtin/portal/spawner.go` and adding it to `builtin/portal/default.go`.
- Add a new builtin module by mirroring `builtin/time/module.go` and mapping in `builtin/module.go`.
- Add a CLI command by mirroring `cmd/cli/run/cli.go` and wiring in `cmd/cli/cli/main.go`.
- Add a REPL command by mirroring `builder/repl/core/command/mutate_node.go` and handling in `builder/repl/core/executor/main.go`.
- Add rule parsing features via `parser/file/rule.go` and `parser/pattern/param_part.go`.
- Add AST nodes via `syntaxtree/rule.go` and `syntaxtree/result/object.go`.
- Add mutation builders in `core/mutation/rule/builder/main.go` and `core/mutation/rule/builder/object.go`.
- Add extractor logic in `core/mutation/rule/extractor/main.go` and `core/mutation/rule/extractor/top_level.go`.
- Add new portal wiring in `builder/program/main.go` and `builder/repl/main.go`.
- Extend list/string/number ops in `builtin/global/list_ops.go`, `builtin/global/string_ops.go`, and `builtin/global/number_ops.go`.

## Tests to scan
- Spawn rule test is in `builtin/global/spawn_test.go`.
- Global control tests are in `builtin/global/control_test.go`.
- Global list ops tests are in `builtin/global/list_ops_test.go`.
- Global number ops tests are in `builtin/global/number_ops_test.go`.
- Global string ops tests are in `builtin/global/string_ops_test.go`.
- Global object ops tests are in `builtin/global/object_ops_test.go`.
- Core base compare tests are in `core/base/compare_test.go`.
- Rule builder structure tests are in `core/mutation/rule/builder/structure_test.go`.
- Parser file module tests are in `parser/file/module_test.go`.
- Parser result structure tests are in `parser/result/structure_test.go`.
- User-defined module build tests are in `core/module/user_defined_test.go`.

## Gotchas
- User-defined modules support at most one file and panic otherwise in `core/module/user_defined.go`.
- `NewCompoundObject` panics if the head is itself an object in `core/base/compound_object.go`.
- Portal `.call` requires exactly one param and string key in `builtin/global/portal.go`.
- `(spawn!)` does not mutate because the rule requires at least one param in `builtin/global/spawn.go`.
- `spawn!` always returns `$` because the spawner returns `base.Null()` in `builtin/portal/spawner.go` and `core/base/convention.go`.
- `sleep` expects a number in `builtin/time/sleep.go`.
- `try` returns `.empty` when child mutation fails in `builtin/global/object_ops.go`.
- Mutation loop halts when `Mutate()` returns empty in `program/main.go`.
- Structure `get`/`set` only works with tag heads in `core/base/structure.go` and `core/base/convention.go`.
- Active rules have priority over normal rules in `core/base/rule_based_class.go`.

## Staleness watchlist
- CLI command list/flags in `cmd/cli/cli/main.go` and `cmd/cli/run/cli.go`.
- Makefile targets and PATH handling in `Makefile` and `shellstartup.sh`.
- Portal key map in `builtin/portal/default.go` and call sites in `builtin/global/stdio.go` and `builtin/global/spawn.go`.
- Global builtin rule list in `builtin/global/code.go` and `builtin/global/mutator.go`.
- Program build wiring in `builder/program/main.go` and `builder/repl/main.go`.
- Mutation loop behavior in `program/main.go` and `core/base/rule_based_class.go`.
- Parser grammar changes in `parser/file/rule.go`, `parser/result/object.go`, and `parser/pattern/param_part.go`.
- REPL command set in `builder/repl/core/command/parser.go` and `builder/repl/core/executor/main.go`.
- Builtin module map in `builtin/module.go` and module implementations like `builtin/time/module.go`.
- Example program references in `examples/tictactoe.mu` and `Makefile`.
