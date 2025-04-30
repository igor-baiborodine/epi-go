package strings

import (
	"strconv"
	"strings"
)

// LookAndSay takes an integer n as input and returns the nth term in the Look-and-Say sequence.
// Parameters:
//
//	n: an integer representing the position in the sequence (must be positive)
//
// Returns:
//
//	the nth term in the Look-and-Say sequence as a string
//
// Time Complexity: O(n * 2^n) - each term can be up to twice the size of the previous term
// Space Complexity: O(2^n) - the length of the result grows exponentially
func LookAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = nextNumber(s)
	}
	return s
}

// nextNumber takes a string representing the current term in the Look-and-Say sequence
// and returns the next term by counting and describing the digits in the input string.
// Parameters:
//
//	s: a string representing the current term in the sequence
//
// Returns:
//
//	the next term in the Look-and-Say sequence as a string
//
// Time Complexity: O(n) where n is the length of input string
// Space Complexity: O(n) for storing the result in strings.Builder
func nextNumber(s string) string {
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			i++
			count++
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(s[i])
	}
	return sb.String()
}
