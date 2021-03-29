# Extending the interpreter and its interactive environment

 _to be continued..._

## Step 3: Add settings `level (program|statement|expression)` and `process (parse|eval|type)`

Every ast node is either a program, a statement or an expression. Until now, we treat each input as a program, which means, we can also only evaluate a program and show the type for the evaluation result of a program.
In this step, we implement another setting that chooses to parse and thus further evaluate the input as either program,statement or expression.
In addition, the commands `expr[ession]`, `stmt|statement` and `prog[ram]` are implemented.

Furthermore, we implement settings for the way the input is to be processed: `(parse|eval|type)`. The commands `type` and `eval` are already implemented, so only `parse` is added.

The full instruction set is now:

| NAME           |                   | USAGE                                                    |
--- | --- | --- |
| clear          | ~                 | clear the environment                                    |
| e[val]         | ~ `<input>`         | print out value of object `<input>` evaluates to           |
| expr[ession]   | ~ `<input>`         | expect `<input>` to be an expression                       |
| h[elp]         | ~                 | list all commands with usage                             |
|                | ~ `<cmd>`           | print usage command `<cmd>`                                |
| l[ist]         | ~                 | list all identifiers in the environment alphabetically   |
|                |                   |      with types and values                               |
| p[arse]        | ~ `<input>`         | parse `<input>`                                            |
| paste          | ~ `<input>`         | evaluate multiline `<input>` (terminated by blank line)    |
| prog[ram]      | ~ `<input>`         | expect `<input>` to be a program                           |
| q[uit]         | ~                 | quit the session                                         |
| reset          | ~ process         | set process to default                                   |
|                | ~ level           | set level to default                                     |
|                | ~ logtype         | set logtype to default                                   |
|                | ~ paste           | set multiline support to default                         |
|                | ~ prompt          | set prompt to default                                    |
| set            | ~ process `<p>`     | `<p>` must be: [e]val, [p]arse, [t]ype                     |
|                | ~ level `<l>`       | `<l>` must be: [p]rogram, [s]tatement, [e]xpression        |
|                | ~ logtype         | when eval, additionally output objecttype                |
|                | ~ paste           | enable multiline support                                 |
|                | ~ prompt `<prompt>` | set prompt string to `<prompt>`                            |
| settings       | ~                 | list all settings with their current values and defaults |
| stmt|statement | ~ `<input>`         | expect `<input>` to be a statement                         |
| t[ype]         | ~ `<input>`         | show objecttype `<input>` evaluates to                     |
| unset          | ~ logtype         | when eval, don't additionally output objecttype          |
|                | ~ paste           | disable multiline support                                |


## Step 2: Implement a small initial instruction set for the interactive environment

- new interactive environment in directory `session`
    - inspo from [ghci](https://downloads.haskell.org/~ghc/latest/docs/html/users_guide/ghci.html#ghci-commands) and [gore](https://github.com/motemen/gore) 

![Demo5](demos/demo5.gif)


| NAME   |                   | USAGE                                                   |
|--------|-------------------|---------------------------------------------------------|
| clear    | ~                 | clear the environment                                    |
| e[val]   | ~ `<input>`         | print out value of object `<input>` evaluates to           |
| h[elp]   | ~                 | list all commands with usage                             |
|          | ~ `<cmd> `          | print usage command `<cmd>`                                |
| l[ist]   | ~                 | list all identifiers in the environment alphabetically   |
|          |                   |      with types and values                               |
| paste    | ~ `<input>`         | evaluate multiline `<input>` (terminated by blank line)    |
| q[uit]   | ~                 | quit the session                                         |
| reset    | ~ logtype         | set logtype to default                                   |
|          | ~ paste           | set multiline support to default                         |
|          | ~ prompt          | set prompt to default                                    |
| set      | ~ logtype         | when eval, additionally output objecttype                |
|          | ~ paste           | enable multiline support                                 |
|          | ~ prompt `<prompt>` | set prompt string to `<prompt> `                           |
| settings | ~                 | list all settings with their current values and defaults |
| t[ype]   | ~ `<input>`         | show objecttype `<input>` evaluates to                     |
| unset    | ~ logtype         | when eval, don't additionally output objecttype          |
|          | ~ paste           | disable multiline support                                |


- if `input` is not prefixed by `:<cmd>`, it is equivalent to `:eval input`


## Step 1: Write Tests for bugs in parser and evaluator

- `parser/parser_add_test.go`
- `evaluator/evaluator_add_test.go`

## Step 0: Starting Point: Copy the Code

- unaltered code from the book [_Writing an Interpreter in Go_](https://interpreterbook.com/), Version 1.7





