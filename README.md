# mutO language
**mutO** (**mut**ating **O**bject) is a programming language that introduces a new paradigm: Everything is Object, unlike the object you know, because they are... (wait for it) ... **MUTATING!** 🎉🎉🎉

- [1. Concept](#section-1)
- [2. Install](#section-2)
- [3. CLI](#section-3)
- [4. Examples](#section-4)

<a name="section-1"></a>
## 1. Concept
### 1.1 Object
Everything is Object. So, what I mean by "Object"? Here is how it looks like:
```muto
f (g 123) "abc"
```
The line above represents a tree with the root node consisting of **head**: *f*; and **children**: *(g 123)*, and *"abc"*. And its children are Object by themselves as well. That's why I said it's a tree!

Note that the head could be an Object as well, eg. `(f "hello") (g (h 123)) "abc"`, in this case, the head is the Object `f "hello"`.

When executed, the Object will mutate itself (along with their children) until they reach **stable state**. That's where the program stops.

### 1.2 Mutation
Mutation is the key mechanism that gradually **changes** our initial Object toward the final result.

When all mutations are no longer be able to apply to the Object, we would say that it reaches its **stable state**. If the root Object is stable, we call it the final result which will be returned to the user, and the program terminates.

There are 3 types of mutations: *Normal Mutation*, *Bubble Up*, and *Active Mutation*. Given an Object *obj*, how program will apply mutations to it can be explained by this pseudo-code:
```
fn mutateOnce(obj):
    case obj can activeMutate:
        return (activeMutate obj)
    case (head obj) is nested Object and can mutate:
        (head obj) <- mutateOnce (head obj)
        return obj
    case obj can bubbleUp:
        return (bubbleUp obj)
    case obj has child that can mutate:
        (thatChild obj) <- mutateOnce (thatChild obj)
        return obj
    otherwise:
        return (normalMutate obj)
```

#### Normal Mutation

Normal Mutation is one that can be coded by user as a rule. Here is an example of a normal mutation rule:
```muto
g X = + X 10
```
This rule tells mutO interpreter that whenever it sees an Object with head: *g* with the first child, let's call it with variable *X*, replace the Object with a new Object `+ X 10`.

For example, if we have Object `g 123`, it will be replaced with a new Object, `+ 123 10`, by this rule.

If there are multiple rules can be applied, mutO chooses the first one.

Normal Mutation is the lowest priority mutation. Meaning it will be applied after every other things is applied, and all of it children are stable.

#### Bubble Up
Bubble up is a mechanism that flatten Object with nested head one-level up **if the head is stable**. For example, given an Object `(f 1) "a"`, if there are no further mutations to be applied to `f 1`, the entire Object will be bubbled up to be `f 1 "a"`. Then, mutO will try to mutate it again.

#### Active Mutation
Active Mutation is another one that can be coded by user as a rule. Unlike Normal Mutation, Active Mutation has the highest priority. Meaning that, it doesn't care whether the head or children are stable or Bubbling Up is already applied or not. It tries to match the Object as is. For example:
```muto
main = f (g 10)

g = * 9
f = + 20

@ f (g X) = X
```
The Active Mutation rule is the rule starting with `@` sign. In the code above, even if there are normal rules defined for `g 10`. The active rule of `f` matches it first and mutates to the final result: `10`.

Active Mutation is a powerful yet dangerous mechanism. It should be used as little as possible!

<a name="section-2"></a>
## 2. Install

First, install Golang v1.22.7.

Then, build using `Makefile`
```shell
go install github.com/SSripilaipong/muto@v0.0.1
```
Make sure your PATH includes GOPATH in rc files such as `~/.bashrc` or `~/.zshrc`:
```shell
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
```

<a name="section-3"></a>
## 3. CLI
### 3.1 run from file
Given file *main.mu*:
```
main = ++ "hello, " who
who = "world"
```
You can execute *main.mu* to get just the final result by:
```shell
$ muto run main.mu
"hello, world"
```
### 3.2 run explain from file
Given file *main.mu*:
```
main = ++ "hello, " who
who = "world"
```
You can let muto explain each mutation steps by:
```shell
$ muto run --explain main.mu
++ "hello, " who
++ "hello, " "world"
"hello, world"
```

<a name="section-4"></a>
## 4. Examples
### Hello World
```muto
main = "Hello, World"
```
Run explain:
```muto
"Hello, World"
```
Note string (and number) is a stable object, so no further mutation occurs.

### Adding numbers
```muto
main = + 1 2
```
Run explain:
```muto
+ 1 2
3
```
`+ 1 2` is an unstable Object, since there are builtin rule for pattern `+ A B`. So it mutates once and become `3`.

### Sum numbers
```muto
main = sum 1 2 3 4

sum X Y = sum (+ X Y)
sum Z = Z
```
Run explain:
```muto
sum 1 2 3 4
sum (+ 1 2) 3 4
sum 3 3 4
sum (+ 3 3) 4
sum 6 4
sum (+ 6 4)
sum 10
10
```
First, pattern `sum X Y` partially matches Object `sum 1 2` and partially mutates to `sum 3`, so the entire Object becomes `sum 3 3 4`

The process repeats until the Object becomes `sum 10` which is matched by pattern `sum Z`, so it mutates to `10`.

### Append children
```muto
append ($ Xs...) ($ Ys...) = $ Xs... Ys...

main = append ($ 1 2) ($ 3 4)
```
Run explain:
```muto
append ($ 1 2) ($ 3 4)
$ 1 2 3 4
```
`Xs...` and `Ys...` are **variadic pattern** which can be expanded in to right-hand side.

Note that `$` is just a name of an Object. You can replace it with any other names you like. Conventionally, `$` is a name for data Object.

### Auto-reducible data type

For example, Optional value of value (T<sup>2</sup>) could be automatically reduced to just Optional value (T):

```muto
main = f' (f 5)

f 0 = opt'empty
f N = opt'value (/ 100 N)

f' = opt'fmap f

opt'fmap F (opt'value X) = opt'value (F X)
opt'fmap F opt'empty = opt'empty

opt'value (opt'value X) = opt'value X
```
Run explain:
```muto
f' (f 5)
f' (opt'value (/ 100 5))
f' (opt'value 20)
opt'fmap f (opt'value 20)
opt'value (f 20)
opt'value (opt'value (/ 100 20))
opt'value (opt'value 5)
opt'value 5
```

### Map
```muto
main = map (* 10) $ ($ 1 2 3)

map F ($ Ys...) ($ X Xs...) = map F ($ Ys... (F X)) ($ Xs...)
map F R $ = R
```
Run explain:
```muto
map (* 10) $ ($ 1 2 3)
map (* 10) ($ ((* 10) 1)) ($ 2 3)
map (* 10) ($ (* 10 1)) ($ 2 3)
map (* 10) ($ 10) ($ 2 3)
map (* 10) ($ 10 ((* 10) 2)) ($ 3)
map (* 10) ($ 10 (* 10 2)) ($ 3)
map (* 10) ($ 10 20) ($ 3)
map (* 10) ($ 10 20 ((* 10) 3)) $
map (* 10) ($ 10 20 (* 10 3)) $
map (* 10) ($ 10 20 30) $
$ 10 20 30
```

### Functional composition
```muto
main = f 2
f = . g h

g = + 10
h = * 16

. F G X = F (G X)
```
Run explain:
```muto
f 2
. g h 2
. (+ 10) h 2
. (+ 10) (* 16) 2
(+ 10) ((* 16) 2)
+ 10 ((* 16) 2)
+ 10 (* 16 2)
+ 10 32
42
```
### Functional composition with unstable Object
Remember `map` from the previous example? If you try to apply the functional composition to it you'll find it not work. Since the `map F R $` is in an **unstable form** which means it ready to mutate. So you can't pass it around like in the previous example.
```muto
map F ($ Ys...) ($ X Xs...) = map F ($ Ys... (F X)) ($ Xs...)
map F R $ = R
```
One way to go is to make the initial Object be a **stable form**. So it doesn't mutate until it meets the first input.
```muto
map F S = map' F $ S

map' F ($ Ys...) ($ X Xs...) = map' F ($ Ys... (F X)) ($ Xs...)
map' F R $ = R
```
Now you can compose it like other Objects:
```muto
main = . (map (+ 10)) (map (* 5)) ($ 1 2 3)
```
Run explain:
```muto
. (map (+ 10)) (map (* 5)) ($ 1 2 3)
(map (+ 10)) ((map (* 5)) ($ 1 2 3))
map (+ 10) ((map (* 5)) ($ 1 2 3))
map (+ 10) (map (* 5) ($ 1 2 3))
map (+ 10) (map' (* 5) $ ($ 1 2 3))
map (+ 10) (map' (* 5) ($ ((* 5) 1)) ($ 2 3))
map (+ 10) (map' (* 5) ($ (* 5 1)) ($ 2 3))
map (+ 10) (map' (* 5) ($ 5) ($ 2 3))
map (+ 10) (map' (* 5) ($ 5 ((* 5) 2)) ($ 3))
map (+ 10) (map' (* 5) ($ 5 (* 5 2)) ($ 3))
map (+ 10) (map' (* 5) ($ 5 10) ($ 3))
map (+ 10) (map' (* 5) ($ 5 10 ((* 5) 3)) $)
map (+ 10) (map' (* 5) ($ 5 10 (* 5 3)) $)
map (+ 10) (map' (* 5) ($ 5 10 15) $)
map (+ 10) (($ 5 10 15))
map (+ 10) ($ 5 10 15)
map' (+ 10) $ ($ 5 10 15)
map' (+ 10) ($ ((+ 10) 5)) ($ 10 15)
map' (+ 10) ($ (+ 10 5)) ($ 10 15)
map' (+ 10) ($ 15) ($ 10 15)
map' (+ 10) ($ 15 ((+ 10) 10)) ($ 15)
map' (+ 10) ($ 15 (+ 10 10)) ($ 15)
map' (+ 10) ($ 15 20) ($ 15)
map' (+ 10) ($ 15 20 ((+ 10) 15)) $
map' (+ 10) ($ 15 20 (+ 10 15)) $
map' (+ 10) ($ 15 20 25) $
$ 15 20 25
```
