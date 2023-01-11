### (WIP) Elements of Programming Interviews in Go

Interview problems and solutions presented in this repository are based on
the [Elements of Programming Interviews](https://elementsofprogramminginterviews.com/).

**Useful resources**

- [Go Programming Language Documentation Portal](https://go.dev/doc/)
- [Go Programming Language Specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Frequently Asked Questions (FAQ)](https://go.dev/doc/faq)
- [Go Wiki](https://github.com/golang/go/wiki)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Primitive Types](#primitive-types)
- [Arrays](#arrays)
- [Strings](#strings)
- [Linked Lists](#linked-lists)
- [Stacks and Queues](#stacks-and-queues)
- [Binary Trees](#binary-trees)
- [Heaps](#heaps)
- [Searching](#searching)
- [Hash Tables](#hash-tables)
- [Sorting](#sorting)
- [Binary Search Trees](#binary-search-trees)
- [Recursion](#recursion)
- [Dynamic Programming](#dynamic-programming)
- [Greedy Algorithms and Invariants](#greedy-algorithms-and-invariants)
- [Graphs](#graphs)
- [Parallel Computing](#parallel-computing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### Primitive Types

**Useful Resources**

- [Bitwise Operation](https://en.wikipedia.org/wiki/Bitwise_operation)
- [Bit Hacking with Go](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)
- [Bit Twiddling Hacks](https://graphics.stanford.edu/~seander/bithacks.html)

**Know Your Primitive Types**

You should know very well Go's basic types, namely, [boolean](https://go.dev/ref/spec#Boolean_types), 
[numeric](https://go.dev/ref/spec#Numeric_types), and [string](https://go.dev/ref/spec#String_types).
You should also know the utility methods for numeric types in [math](https://pkg.go.dev/math) package.

- Be very familiar with the bit-wise operators such as `6&4`,  `1|2`, `8>>1`, `-16>>2`, `1<<10`, `15^10`, `4&^3` and key
  methods in the [math/bits](https://pkg.go.dev/math/bits) package.
- Know the constants from the `math` package denoting the maximum and minimum values for numeric types,
  e.g., `math.MaxInt`, `math.MaxFloat32`, `math.MinInt`.
- The key methods for numeric types are `math.Abs(-2)`, `math.Ceil(1.49)`, `math.Floor(1.51)`, `math.Min(-2, -4)`, `math.Max(3.14, 5.16)`, `math.Pow(2, 3)`, and `math.Sqrt(25)`.
- Know about explicit type conversion between different numeric types, e.g. `float64(42)`, `byte(20)`, numeric types
  to `string`, e.g., `strconv.Itoa(12)`, `string` to numeric types, e.g., `strconv.Atoi("41")`.
- Be familiar with the key functions in the [math/rand](https://pkg.go.dev/math/rand) package.

**Tasks**

- [X] [Computing the parity of a word](/primitivetypes/parity.go), [tests](/primitivetypes/parity_test.go)
- [X] [Swap bits](/primitivetypes/swapbits.go), [tests](/primitivetypes/swapbits_test.go)
- [X] [Reverse bits](/primitivetypes/reversebits.go), [tests](/primitivetypes/reversebits_test.go)
- [ ] Find the closest integer with the same weight
- [ ] Compute `x*y` without arithmetical operators
- [ ] Compute `x/y`
- [ ] Compute `x**y`
- [X] [Reverse digits](/primitivetypes/reversedigits.go), [tests](/primitivetypes/reversebits_test.go)
- [ ] Check if a decimal integer is a palindrome
- [ ] Generate uniform random numbers
- [ ] Rectangle intersection

### Arrays

**Useful Resources**

- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Slice Tricks](https://github.com/golang/go/wiki/SliceTricks)

**Know Your Array Libraries**

The base [array](https://go.dev/ref/spec#Array_types) type in Go has a fixed size, and its length is part of its type.
Arrays are rarely used directly in Go due to their rigidity. On the other hand,
the [slice](https://go.dev/ref/spec#Slice_types) type which is backed by an underlying array is more flexible because
the length is not part of its type. Therefore, slices are more common than arrays. Use arrays when you know the exact
length you need ahead of time.

- Know the syntax for allocating and initializing an array, i.e., `[3]int`, `[3]int{10, 20, 30}`, `[...]int{10, 20, 30}`, `[12]int{1, 5: 4, 6, 10: 100, 15}`.
- Arrays can be compared using `==` and `!=` operators.
- Know how to allocate a `2D` array, e.g., `[2][3]int`.
- Array's length can be obtained using the `len` built-in function.
- Know the syntax for allocating and initializing a slice, e.g., `[]int`, `[]int{}`, `[]int{10, 20, 30}`, `[]int{1, 5: 4, 6, 10: 100, 15}`.
- Know how to allocate a `2D` slice, e.g., `[][]int`.
- Know the [difference](https://go.dev/doc/effective_go#allocation_make) between allocating a slice using the `new`
  and `make` functions.
- Be familiar with slice expressions, e.g., `x[:2]`, `x[1:]`, `x[1:3]`, `x[:]`.  
- Key built-in functions for slices include `len`, `cap`, `append`, `copy`.
- Know how to compare slices using `bytes.Equals` for `[]byte` slices and `reflect.DeepEqual` for testing purposes.
- Be familiar with the functions in the [sort](https://pkg.go.dev/sort) package, e.g., `sort.Ints`, `sort.Float64s`, `sort.Strings`.

**Tasks**

- [X] [The Dutch national flag problem](/arrays/dutchflag.go), [tests](/arrays/dutchflag_test.go)
- [X] [Increment an arbitrary-precision integer](/arrays/plusone.go), [tests](/arrays/plusone_test.go)
- [ ] Multiply two arbitrary-precision integers
- [ ] Advancing through an array
- [X] [Delete duplicates from a sorted array](/arrays/duplicates.go), [tests](/arrays/duplicates_test.go)
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

**Useful Resources**

- [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)

**Know Your String Libraries**

- Know operators for manipulating strings: compare(`==`, `!=`), order(`>`, `>=`, `<`, `<=`), and concatenate(`+`).
- Important to remember that strings in Go are immutable.
- Know that index expressions to extract a single value, i.g., `s[4]`, and slice expressions to extract substrings, e.g., `s[3:5]`, `s[:8]`, `s[2:]`, can be used only for strings containing characters that take up one byte, e.g., [ASCII](https://en.wikipedia.org/wiki/ASCII) characters.
- Know what code point([rune](https://go.dev/ref/spec#Rune_literals)) is, and how to iterate through code points in a given string using the `range` clause.
- Know how to do conversions between strings, runes, and bytes.
- Be familiar with the functions in the [strings](https://pkg.go.dev/strings) and [unicode/utf8](https://pkg.go.dev/unicode/utf8) packages, e.g., `strings.ToUpper`, `strings.Contains`, `strings.Join`, `strings.Split`, `strings.ReplaceAll`.
  
**Tasks**

- [ ] Interconvert strings and integers
- [ ] Base conversion
- [ ] Compute the spreadsheet column encoding
- [X] [Replace and remove](/strings/replaceremove.go), [tests](/strings/replaceremove_test.go)
- [X] [Test palindromicity](/strings/palindromicity.go), [tests](/strings/palindromicity_test.go)
- [ ] Reverse all the words in a sentence
- [ ] Compute all mnemonics for a phone number
- [ ] The look-and-say problem
- [ ] Convert from Roman to decimal
- [X] [Compute all valid IP addresses](/strings/ipaddress.go), [tests](/strings/ipaddress_test.go)
- [ ] Write a string sinusoidally
- [ ] Implement run-length encoding
- [ ] Find the first occurrence of a substring 

### Linked Lists
WIP

**Know Your Linked Lists Libraries**

**Tasks**

- [ ] Merge two sorted lists
- [ ] Reverse a single sublist
- [ ] Test for cyclicity
- [ ] Test for overlapping lists - lists are cycle-free
- [ ] Test for overlapping lists - lists may have cycles
- [ ] Delete a node from a singly linked list
- [ ] Remove the kth last element from a list
- [ ] Remove duplicates from a sorted list
- [ ] Implement cyclic right shift for singly linked lists
- [ ] Implement even-odd merge
- [ ] Test whether a singly linked list is palindromic
- [ ] Implement list pivoting
- [ ] Add list-based integers

### Stacks and Queues
WIP

**Know Your Queue Libraries**

**Know Your Stack Libraries**

**Tasks**

- [ ] Implement a stack with max API
- [ ] Evaluate RPN expressions
- [ ] Test a string over `{,},(,),[,]` for well-formedness
- [ ] Normalize path names
- [ ] Compute buildings with a sunset view
- [ ] Compute binary tree nodes in order of increasing depth
- [ ] Implement a circular queue
- [ ] Implement a queue using stacks
- [ ] Implement a queue with max API

### Binary Trees
WIP

**Tasks**

- [ ] Test if a binary tree is height-balanced
- [ ] Test if a binary tree is symmetric
- [ ] Compute the lowest common ancestor in a binary tree
- [ ] Compute the LCA when nodes have parent pointers
- [ ] Sum the root-to-leaf paths in a binary tree
- [ ] Find a root to leaf path with specified sum
- [ ] Implement an inorder traversal without recursion
- [ ] Implement a preorder traversal without recursion
- [ ] Compute the ***k***th node in an inorder traversal
- [ ] Compute the successor
- [ ] Implement an inorder traversal with ***O***(1) space .
- [ ] Reconstruct a binary tree from traversal data
- [ ] Reconstruct a binary tree from a preorder traversal with markers
- [ ] Form a linked list from the leaves of a binary tree .
- [ ] Compute the exterior of a binary tree .
- [ ] Compute the right sibling tree .

### Heaps
WIP

**Know Your Heap Libraries**

**Tasks**

- [ ] Merge sorted files
- [ ] Sort an increasing-decreasing array
- [ ] Sort an almost-sorted array
- [ ] Compute the ***k*** closest stars
- [ ] Compute the median of online data .
- [ ] Compute the ***k*** largest elements in a max-heap

### Searching
WIP

**Know Your Searching Libraries**

**Tasks**

- [ ] Search a sorted array for first occurrence of ***k***
- [ ] Search a sorted array for entry equal to its index
- [ ] Search a cyclically sorted array
- [ ] Compute the integer square root
- [ ] Compute the real square root
- [ ] Search in a `2D` sorted array
- [ ] Find the min and max simultaneously
- [ ] Find the ***k***th largest element
- [ ] Find the missing IP address
- [ ] Find the duplicate and missing elements

### Hash Tables
WIP

**Know Your Hash Tables Libraries**

**Tasks**
- [ ] Test for palindromic permutations
- [ ] Is an anonymous letter constructible?
- [ ] Implement an ISBN cache
- [ ] Compute the LCA, optimizing for close ancestors
- [ ] Find the nearest repeated entries in an array
- [ ] Find the smallest subarray covering all values
- [ ] Find the smallest subarray sequentially covering all values
- [ ] Find the longest subarray with distinct entries
- [ ] Find the length of a longest contained interval
- [ ] Compute all string decompositions
- [ ] Test the Collatz conjecture
- [ ] Implement a hash function for chess

### Sorting
WIP

**Know Your Sorting Libraries**

**Tasks**

- [ ] Compute the intersection of two sorted arrays
- [ ] Merge two sorted arrays
- [ ] Remove first-name duplicates
- [ ] Smallest non-constructible value
- [ ] Render a calendar
- [ ] Merging intervals
- [ ] Compute the union of intervals
- [ ] Partitioning and sorting an array with many repeated entries
- [ ] Team photo day-1
- [ ] Implement a fast sorting algorithm for lists
- [ ] Compute a salary threshold

### Binary Search Trees
WIP

**Know Your Binary Search Tree Libraries**

**Tasks**

- [ ] Test if a binary tree satisfies the BST property
- [ ] Find the first key greater than a given value in a BST
- [ ] Find the ***k*** largest elements in a BST
- [ ] Compute the LCA in a BST
- [ ] Reconstruct a BST from traversal data
- [ ] Find the closest entries in three sorted arrays
- [ ] Enumerate numbers of the form ***a*** + ***b***√2
- [ ] Build a minimum height BST from a sorted array
- [ ] Test if three BST nodes are totally ordered
- [ ] The range lookup problem
- [ ] Add credits

### Recursion
WIP

**Know Your Recursion Libraries**

**Tasks**

- [ ] The Towers of Hanoi problem
- [ ] Generate all non-attacking placements of ***n***-Queens
- [ ] Generate permutations
- [ ] Generate the power set
- [ ] Generate all subsets of size ***k***
- [ ] Generate strings of matched parens
- [ ] Generate palindromic decompositions
- [ ] Generate binary trees
- [ ] Implement a Sudoku solver
- [ ] Compute a Gray code

### Dynamic Programming
WIP

**Know Your Dynamic Programming Libraries**

**Tasks**

- [ ] Count the number of score combinations
- [ ] Compute the Levenshtein distance
- [ ] Count the number of ways to traverse a `2D` array
- [ ] Compute the binomial coefficients
- [ ] Search for a sequence in a `2D` array
- [ ] The knapsack problem
- [ ] The `bedbathandbeyond.com` problem
- [ ] Find the minimum weight path in a triangle
- [ ] Pick up coins for maximum gain
- [ ] Count the number of moves to climb stairs
- [ ] The pretty printing problem
- [ ] Find the longest non-decreasing subsequence

### Greedy Algorithms and Invariants
WIP

**Know Your Greedy Algorithms and Invariants Libraries**

**Tasks**

- [ ] Compute an optimum assignment of tasks
- [ ] Schedule to minimize waiting time
- [ ] The interval covering problem
- [ ] The 3-sum problem
- [ ] Find the majority element
- [ ] The gasup problem
- [ ] Compute the maximum water trapped by a pair of vertical lines
- [ ] Compute the largest rectangle under the skyline
 
### Graphs
WIP

**Know Your Graphs Libraries**

**Tasks**

- [ ] Search amaze
- [ ] Paint a Boolean matrix
- [ ] Compute enclosed regions
- [ ] Deadlock detection
- [ ] Clone a graph
- [ ] Making wired connections
- [ ] Transform one string to another
- [ ] Team photo day-2

### Parallel Computing
WIP

**Know Your Parallel Computing Libraries**

**Tasks**

- [ ] Implement caching for a multi-threaded dictionary
- [ ] Analyze two unsynchronized interleaved threads
- [ ] Implement synchronization for two interleaving threads
- [ ] Implement a thread pool
- [ ] Deadlock
- [ ] The readers-writers problem
- [ ] The readers-writers problem with write preference
- [ ] Implement a Timer class
- [ ] Test the Collatz conjecture in parallel
