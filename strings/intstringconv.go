package strings

import (
	"slices"
	"strings"
)

// IntToString converts an integer to its string representation.
// If the integer is negative, a '-' sign is prefixed in the result.
// The function iterates through each digit of the number, converting it to its
// corresponding character, and builds the resulting string in reverse order using
// a strings.Builder. The helper function reverseString is then used to reverse
// the string to its correct order.
//
// Time Complexity: O(d), where d is the number of digits in the integer n.
// Space Complexity: O(d), additional space is used to store the reversed string.
func IntToString(n int) string {
	var neg bool
	if n < 0 {
		neg = true
		n = -n
	}
	sb := strings.Builder{}

	for n != 0 {
		sb.WriteString(string(rune('0' + n%10)))
		n /= 10
	}
	if neg {
		sb.WriteString("-")
	}
	return reverseString(sb.String())
}

// ReverseString reverses the input string and returns the reversed string.
// It first converts the string into a slice of runes to correctly handle Unicode characters
// (ensuring multibyte characters are reversed properly). Then, it swaps runes iteratively from
// the beginning and the end toward the middle of the slice.
//
// Time Complexity: O(n), where n is the length of the string. Each character is visited once during
// the swap process.
// Space Complexity: O(n), additional space is used to convert the string to a rune slice.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseString2(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func StringToInt(s string) int {
	var result int
	var neg bool

	i := 0
	if s[0] == '-' {
		i = 1
		neg = true
	}
	for _, c := range s[i:] {
		result = result*10 + int(c-'0') // compute the difference between the code point of `c` and `'0'`

	}
	if neg {
		result = -result
	}
	return result
}
