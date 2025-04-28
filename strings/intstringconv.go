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
	return ReverseString(sb.String())
}

// StringToInt converts a string representation of an integer to its numerical form.
// It handles optional negative signs at the beginning of the string.
// The function iterates through each character of the string, computes its numerical value,
// and combines it into the resultant integer.
// If a '-' sign is detected, the result is negated.
//
// Time Complexity: O(n), where n is the length of the string.
// Space Complexity: O(1), as no additional space proportional to input size is used.
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

func ReverseString(s string) string {
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
