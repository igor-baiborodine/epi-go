package strings

import (
	"strings"
	"unicode"
)

// Palindromicity reports whether a string is palindromic,
// given that non-alphanumeric characters are ignored and
// characters are compared case-insensitively.
// The time complexity is O(n).
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
