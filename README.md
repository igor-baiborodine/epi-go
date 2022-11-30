### Elements of Programming Interviews in Go

Interview problems and solutions presented in this repository are based on
the [Elements of Programming Interviews](https://elementsofprogramminginterviews.com/).

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Go Programming Language Documentation Portal](#go-programming-language-documentation-portal)
- [Primitive Types](#primitive-types)
- [Arrays](#arrays)
- [Strings](#strings)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### [Go Programming Language Documentation Portal](https://go.dev/doc/)

### Primitive Types

Useful resources:
- [Bitwise Operation](https://en.wikipedia.org/wiki/Bitwise_operation)
- [Bit Hacking with Go](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)
- [Bit Twiddling Hacks](https://graphics.stanford.edu/~seander/bithacks.html)

You should know very well Go's basic types, namely, [boolean](https://go.dev/ref/spec#Boolean_types)
, [numeric](https://go.dev/ref/spec#Numeric_types), and [string](https://go.dev/ref/spec#String_types). 
You should also know the utility methods for numeric types in [math](https://pkg.go.dev/math) package.
 
- Be very familiar with the bit-wise operators such as `6&4`,  `1|2`, `8>>1`, `-16>>2`, `1<<10`, `15^10`, `4&^3`.
- Know the constants from the `math` package denoting the maximum and minimum values for numeric types,
  e.g., `math.MaxInt`, `math.MaxFloat32`, `math.MinInt`.
- The key methods for numeric types are `math.Abs(-2)`, `math.Ceil(1.49)`, `math.Floor(1.51)`, `math.Min(-2, -4)`
  , `math.Max(3.14, 5.16)`, `math.Pow(2, 3)`, and `math.Sqrt(25)`.
- Know about explicit type conversion between different numeric types, e.g. `float64(42)`, `byte(20)`, numeric types
  to `string`, e.g., `strconv.Itoa(12)`, `string` to numeric types, e.g., `strconv.Atoi("41")`, and `string` to a slice
  of runes, e.g., `[]rune("ABC€")`.
- Be familiar with the key methods in the [math/rand](https://pkg.go.dev/math/rand) package.

**Tasks**
1. [ ] Computing the parity of a word
2. [ ] Swap bits
3. [ ] Reverse bits
4. [ ] Find a closest integer with the same weight
5. [ ] Compute x X y without arithmetical operators
6. [ ] Compute x/y
7. [ ] Compute x ** y
8. [ ] Reverse digits
9. [ ] Check if a decimal integer is a palindrome
10. [ ] Generate uniform random numbers
11. [ ] Rectangle intersection

### Arrays
WIP

### Strings
WIP 
