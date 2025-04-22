package strings

import (
	"strings"
	"unicode"
)

// Palindromicity checks whether a given string s is a palindrome,
// considering only alphanumeric characters and ignoring case.
// Non-alphanumeric characters are skipped during the comparison.
//
// Time Complexity: O(n), where n is the length of the input string.
// Each character is processed at most once as we move two pointers from
// the start and end of the string.
//
// Space Complexity: O(1), as the algorithm uses a constant amount of extra memory.
func Palindromicity(s string) bool {
	var i, j = 0, len(s) - 1

	for i < j {
		if !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
			i++
			continue
		}
		if !(unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j]))) {
			j--
			continue
		}
		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
