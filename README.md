# mutO

mutO (mutating Object) is a rewriting language where programs are rules and evaluation is repeated mutation of a root object.

## Install

```bash
go install github.com/SSripilaipong/muto@latest
```

## Run

```bash
# Run a file
muto run main.mu

# Show each mutation step
muto run --explain main.mu

# Interactive REPL
muto repl
```

## Language

### Objects

An object is a head followed by parameters:

```
(f 1 "hello" (g 2))
```

This is a tree with head `f` and three parameters. Parameters can be nested objects.

### Rules

Rules rewrite objects:

```
add X Y = + X Y
```

When the evaluator sees `add 1 2`, it rewrites to `+ 1 2`, then to `3`.

### Variables

- Uppercase names are variables: `X`, `Y`, `Name`
- `_` prefix ignores the binding: `_X` matches but doesn't bind
- `...` suffix matches multiple values: `Xs...`

```
sum X Y Zs... = sum (+ X Y) Zs...
sum X = X

main = sum 1 2 3 4 5
```

### Mutation order

Each step:
1. Try active rules on the object as-is
2. Otherwise, mutate the leftmost mutable child
3. Otherwise, try normal rules
4. If nothing applies, the object is stable

### Active rules

Active rules (prefixed with `@`) match before children mutate:

```
@ f (g X) = X

main = f (g 10)
```

Result: `10` (the active rule matches `f (g 10)` directly)

Without `@`, the child `(g 10)` would mutate first.

### Pattern conjunctions

Bind the same value to multiple patterns using `^`:

```
f X^(g Y) = $ X Y
```

Matches `f (g 1)` with `X=(g 1)` and `Y=1`.

For determinant (head) conjunctions, bind intermediate objects:

```
(f X)^P Y = $ X P Y
```

Matches `(f 1) 2` with `X=1`, `P=(f 1)`, `Y=2`.

### Tags

Tags start with `.` and are literal symbols:

```
(result .ok) = "success"
(result .error) = "failed"
```

### Structures

Structures are key-value stores using tags:

```
person = { .name: "Alice", .age: 30 }

main = person (.get .name)
```

Result: `"Alice"`

### Lists

Use `$` for conventional lists:

```
nums = $ 1 2 3

first ($ X Xs...) = X

main = first nums
```

Result: `1`

## Builtins

### Arithmetic
| Name | Example | Result |
|------|---------|--------|
| `+` | `+ 1 2` | `3` |
| `-` | `- 5 3` | `2` |
| `*` | `* 4 5` | `20` |
| `/` | `/ 10 3` | `3.333...` |
| `div` | `div 10 3` | `3` |
| `mod` | `mod 10 3` | `1` |

### Comparison
| Name | Example | Result |
|------|---------|--------|
| `==` | `== 5 5` | `true` |
| `>` | `> 5 3` | `true` |
| `>=` | `>= 5 5` | `true` |
| `<` | `< 3 5` | `true` |
| `<=` | `<= 5 5` | `true` |

Works on numbers and strings.

### Boolean
| Name | Example | Result |
|------|---------|--------|
| `&` | `& true false` | `false` |
| `\|` | `\| true false` | `true` |
| `!` | `! true` | `false` |

### String
| Name | Example | Result |
|------|---------|--------|
| `++` | `++ "hello" " world"` | `"hello world"` |
| `string` | `string 42` | `"42"` |
| `string-to-runes` | `string-to-runes "ab"` | `$ 'a' 'b'` |

### Control flow

**match** - Pattern matching:
```
(match
  \1 ["one"]
  \2 ["two"]
  \X [X]
) 3
```
Result: `3`

**do** - Sequence, return the last:
```
do (print! "a") (print! "b") "done"
```

**compose** - Function composition:
```
(compose (+ 1) (* 2)) 3
```
Result: `7` (multiply by 2, then add 1)

**curry** - Partial application:
```
add5 = (curry + 5)
main = add5 3
```
Result: `8`

**with** - Apply arguments to function:
```
(with 1 2) +
```
Result: `3`

### Result types

**ok/error** - Result wrappers:
```
(ok 42) .ok?            -- true
(ok 42) .error?         -- false
(ok 42) .value          -- 42
(error "oops") .ok?     -- false
(error "oops") .error?  -- true
(error "oops") .error   -- "oops"
```

**try** - Catch mutation failure:
```
try f 1 2
```
Returns `.value result` if mutation succeeds, `.empty` if it fails.

### List operations

**map** - Transform each element:
```
map (curry * 2) ($ 1 2 3)
```
Result: `$ 2 4 6`

**filter** - Keep matching elements:
```
filter (curry > 2) ($ 1 2 3 4)
```
Result: `$ 3 4`

## IO

```
-- Print to stdout
print! "hello"

-- Read line from stdin
input!

-- Spawn concurrent mutation
spawn! (long-computation 42)
```

## Modules

Import builtin modules in REPL or files:

```
:import time

main = time.sleep 1
```

## CLI

```bash
muto run <file>           # Run a file
muto run --explain <file> # Show each mutation step
muto repl                 # Start REPL
```

REPL commands:
- Type an expression to evaluate it
- Type a rule to add it
- `:import <module>` to import a module
- `:quit` to exit

## Examples

### Factorial

```
factorial 0 = ret 1
factorial N = * N (factorial (- N 1))

main = factorial 5
```

Result: `120`

### Fibonacci

```
fib 0 = ret 0
fib 1 = ret 1
fib N = + (fib (- N 1)) (fib (- N 2))

main = fib 10
```

Result: `55`

### FizzBuzz

```
fizzbuzz N = (match
  \true  [fb (mod N 3) (mod N 5)]
  \false [ret N]
) (| (== 0 (mod N 3)) (== 0 (mod N 5)))

fb 0 0 = ret "FizzBuzz"
fb 0 _ = ret "Fizz"
fb _ 0 = ret "Buzz"

main = map fizzbuzz ($ 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15)
```
