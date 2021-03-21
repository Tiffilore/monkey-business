
# Testing the parser

**Table of contents**
1. [Existing Tests](#existing)
2. [Additional Tests](#additional)
    1. [TestBlockStatementsParseError](#test1)
    2. [TestFunctionLiteralsInvalidParameterSingle](#test2)
    3. [TestFunctionLiteralsInvalidParameterOutsider](#test3)
    4. [TestFunctionLiteralsInvalidParametersCrowd](#test4)

3. [Asts without the Token fields for some of the test data](#asts_without)
3. [Asts with the Token fields for the same subset of the test data](#asts_with)



## Existing Tests: `parser_test.go` <a name="existing"></a>

All existing tests test only input that is supposed to be valid

## Additional Tests: `parser/`parser_add_test.go` <a name="additional"></a>

The added tests only test certain cases that the parser right now perceives as valid, but shouldn't be valid in my view.

### `TestBlockStatementsParseError` <a name="test1"></a>

- tests cases of missing right braces in block statements:

    - as part of an if expression: `if(1){1`
    - as part of a function literal: `fn(x,y){x+y`

### `TestFunctionLiteralsInvalidParameterSingle` <a name="test2"></a>

- tests cases in which random tokens, even illegal ones **(!)** should not be validated as identifier nodes by the parser. 

- selection:

```
fn(@){}     // ILLEGAL
fn(0){}     // INT
fn(=){}     // ASSIGN
fn(!=){}    // NOT_EQ
fn(,){}     // COMMA
fn(fn){}    // FUNCTION
fn(false){} // FALSE
fn(if){}    // IF
fn()){}     // RBRACE
```

- The parser in its current state already rejects `RBRACE` as Parameter, it is just added for completeness

### `TestFunctionLiteralsInvalidParameterOutsider` <a name="test3"></a>

- tests the same parameters surrounded by valid ones, e.g. `fn(a, @, b){}`

### `TestFunctionLiteralsInvalidParametersCrowd` <a name="test4"></a>

- tests a function literal with a lot of invalid parameters followed by an illegal statement (`@`).
- first tests whether there is more than one error reported ( which is not the case right now, but should be)
- then tests whether the last error is the complaint about the illegal end statement
   - thereby wants to test whether the parser parses the input until the end

## Asts without the Token fields for some of the test data <a name="asts_without"></a>

---

- `if(1){1`

<img src="images/ast_wo_tok0.png" width="600" />

---

- `fn(x,y){x+y`

<img src="images/ast_wo_tok1.png" width="600" />

---

- `fn(@){}`

<img src="images/ast_wo_tok2.png" width="600" />

---

- `fn(0){}`

<img src="images/ast_wo_tok3.png" width="600" />

---

- `fn(=){}`

<img src="images/ast_wo_tok4.png" width="600" />

---

- `fn(!=){}`

<img src="images/ast_wo_tok5.png" width="600" />
   
---

- `fn(,){}`

<img src="images/ast_wo_tok6.png" width="600" />

---

- `fn(fn){}`

<img src="images/ast_wo_tok7.png" width="600" />

---

- `fn(false){}`

<img src="images/ast_wo_tok8.png" width="600" />

---

- `fn(if){}`

<img src="images/ast_wo_tok9.png" width="600" />
 
---


## Asts with the Token fields for the same subset of the test data <a name="asts_with"></a>

---

- `if(1){1`

<img src="images/ast_with_tok0.png" width="800" />

---

- `fn(x,y){x+y`

<img src="images/ast_with_tok1.png" width="800" />

---

- `fn(@){}`

<img src="images/ast_with_tok2.png" width="800" />

---

- `fn(0){}`

<img src="images/ast_with_tok3.png" width="800" />

---

- `fn(=){}`

<img src="images/ast_with_tok4.png" width="800" />

---

- `fn(!=){}`

<img src="images/ast_with_tok5.png" width="800" />
   
---

- `fn(,){}`

<img src="images/ast_with_tok6.png" width="800" />

---

- `fn(fn){}`

<img src="images/ast_with_tok7.png" width="800" />

---

- `fn(false){}`

<img src="images/ast_with_tok8.png" width="800" />

---

- `fn(if){}`

<img src="images/ast_with_tok9.png" width="800" />
 
---