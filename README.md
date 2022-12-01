### Elements of Programming Interviews in Go

Interview problems and solutions presented in this repository are based on
the [Elements of Programming Interviews](https://elementsofprogramminginterviews.com/).

**Useful resources**
- [Go Programming Language Documentation Portal](https://go.dev/doc/)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Primitive Types](#primitive-types)
- [Arrays](#arrays)
- [Strings](#strings)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### Primitive Types

**Useful resources**
- [Bitwise Operation](https://en.wikipedia.org/wiki/Bitwise_operation)
- [Bit Hacking with Go](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)
- [Bit Twiddling Hacks](https://graphics.stanford.edu/~seander/bithacks.html)

You should know very well Go's basic types, namely, [boolean](https://go.dev/ref/spec#Boolean_types), 
[numeric](https://go.dev/ref/spec#Numeric_types), and [string](https://go.dev/ref/spec#String_types).
You should also know the utility methods for numeric types in [math](https://pkg.go.dev/math) package.

- Be very familiar with the bit-wise operators such as `6&4`,  `1|2`, `8>>1`, `-16>>2`, `1<<10`, `15^10`, `4&^3` and key
  methods in the [math/bits](https://pkg.go.dev/math/bits) package.
- Know the constants from the `math` package denoting the maximum and minimum values for numeric types,
  e.g., `math.MaxInt`, `math.MaxFloat32`, `math.MinInt`.
- The key methods for numeric types are `math.Abs(-2)`, `math.Ceil(1.49)`, `math.Floor(1.51)`, `math.Min(-2, -4)`
  , `math.Max(3.14, 5.16)`, `math.Pow(2, 3)`, and `math.Sqrt(25)`.
- Know about explicit type conversion between different numeric types, e.g. `float64(42)`, `byte(20)`, numeric types
  to `string`, e.g., `strconv.Itoa(12)`, `string` to numeric types, e.g., `strconv.Atoi("41")`.
- Be familiar with the key methods in the [math/rand](https://pkg.go.dev/math/rand) package.
  
**Tasks**
- [X] [Computing the parity of a word](/primitivetypes/parity.go)
- [X] [Swap bits](/primitivetypes/swapbits.go)
- [X] [Reverse bits](/primitivetypes/reversebits.go)
- [ ] Find the closest integer with the same weight
- [ ] Compute `x*y` without arithmetical operators
- [ ] Compute `x/y`
- [ ] Compute `x**y`
- [ ] Reverse digits
- [ ] Check if a decimal integer is a palindrome
- [ ] Generate uniform random numbers
- [ ] Rectangle intersection

### Arrays
- [ ] The Dutch national flag problem
- [ ] Increment an arbitrary-precision integer
- [ ] Multiply two arbitrary-precision integers
- [ ] Advancing through an array
- [ ] Delete duplicates from a sorted array
- [ ] Buy and sell a stock once
- [ ] Buy and sell a stock twice
- [ ] Enumerate all primes to `n`
- [ ] Permute the elements of an array
- [ ] Compute the next permutation
- [ ] Sample offline data
- [ ] Sample online data
- [ ] Compute a random permutation
- [ ] Compute a random subset
- [ ] Generate nonuniform random numbers
- [ ] The Sudoku checker problem
- [ ] Compute the spiral ordering of a 2D array
- [ ] Rotate a 2D array
- [ ] Compute rows in Pascal's Triangle

### Strings
- [ ] Interconvert strings and integers
- [ ] Base conversion
- [ ] Compute the spreadsheet column encoding
- [ ] Replace and remove
- [ ] Test palindromicity
- [ ] Reverse all the words in a sentence
- [ ] Compute all mnemonics for a phone number
- [ ] The look-and-say problem
- [ ] Convert from Roman to decimal
- [ ] Compute all valid IP addresses
- [ ] Write a string sinusoidally
- [ ] Implement run-length encoding
- [ ] Find the first occurrence of a substring 
