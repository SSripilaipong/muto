# mutO language

mutO (mutating Object) is a small rewriting language where programs are a set of rules and evaluation is repeated mutation of a single root object. Objects are trees: a head plus zero or more parameters, and every parameter is itself an object.

## Quick start

1) Install Go 1.23.8 or newer.
2) Install the CLI:

```shell
go install github.com/SSripilaipong/muto@latest
```

3) Run a file:

```shell
muto run main.mu
```

4) Explain each mutation step:

```shell
muto run --explain main.mu
```

5) Start the REPL:

```shell
muto repl
```

In the REPL, import builtin modules with `:import` before use:

```text
:import time
time.sleep 0
```

If you are working from source:

```shell
make build
./build/muto run examples/tictactoe.mu
```

## How the language works

### Objects
An object is written as a head followed by parameters. For example:

```muto
f (g 123) "abc"
```

This is a tree where the head is `f` and the parameters are `(g 123)` and `"abc"`. Parameters can be objects too.

### Rules
Rules rewrite objects. A rule is written as:

```muto
name pattern = result
```

Example:

```muto
g X = + X 10
```

When the evaluator sees `g 123`, it becomes `+ 123 10`. If multiple rules can apply, the first matching rule is used.

### Mutation order
Each mutation step for a normal class follows this order:

1) Try active rules on the current object.
2) Otherwise, mutate the first mutable child (left-to-right across parameter groups).
3) Otherwise, try normal rules.
4) If nothing applies, the object is stable.

### Active rules
Active rules start with `@` and have higher priority than normal rules. They match the object as-is, before children are mutated.

```muto
main = f (g 10)

g = * 9
f = + 20

@ f (g X) = X
```

The active rule matches `f (g 10)` first, so the result is `10`.

### Program execution
A program starts by evaluating the `main` object. The evaluator keeps mutating until the object is stable, then prints the final result. With `--explain`, every mutation result is printed.

## Syntax tips

- Variables are usually uppercase: `X`, `Y`, `Xs...`.
- Variadic patterns use `...`:

```muto
append ($ Xs...) ($ Ys...) = $ Xs... Ys...
```

- `$` is just a normal name, but is commonly used for data objects.

## Examples

### Hello world

```muto
main = "Hello, World"
```

Explain:

```text
"Hello, World"
```

### Adding numbers

```muto
main = + 1 2
```

Explain:

```text
+ 1 2
3
```

### Summation

```muto
main = sum 1 2 3 4

sum X Y = sum (+ X Y)
sum Z = Z
```

Explain (shortened):

```text
sum 1 2 3 4
sum (+ 1 2) 3 4
...
10
```

### Map

```muto
main = map (* 10) $ ($ 1 2 3)

map F ($ Ys...) ($ X Xs...) = map F ($ Ys... (F X)) ($ Xs...)
map F R $ = R
```

## IO and modules

Builtins include arithmetic, strings, lists, and object utilities. IO is provided via `print!`, `input!`, and `spawn!`.

To use the time module:

```muto
:import time

main = time.sleep 1
```

## CLI reference

- `muto run <file>`: run a file
- `muto run --explain <file>`: print each mutation step
- `muto repl`: start the interactive REPL
