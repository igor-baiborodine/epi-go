package strings

import (
	"strconv"
	"strings"
	"unicode"
)

// EncodeRunLength compresses a string using run-length encoding.
// Parameters:
//   - s: input string to be compressed
//
// Returns:
//   - string: compressed string where each character is preceded by its count
//
// Time Complexity: O(n), where n is the length of the input string
// Space Complexity: O(n) for the output string builder
func EncodeRunLength(s string) string {
	r := []rune(s)
	if len(r) == 1 {
		return "1" + s
	}
	sb := strings.Builder{}
	count := 1

	for i := 1; i <= len(r); i++ {
		if i == len(r) || r[i] != r[i-1] {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteString(string(r[i-1]))
			count = 1
		} else {
			count++
		}
	}
	return sb.String()
}

// DecodeRunLength decompresses a run-length encoded string.
// Parameters:
//   - s: input string in run-length encoded format
//
// Returns:
//   - string: decompressed string where repeated characters are expanded
//
// Time Complexity: O(n), where n is the length of the decompressed string
// Space Complexity: O(n) for the output string builder
func DecodeRunLength(s string) string {
	r := []rune(s)
	sb := strings.Builder{}
	count := 0

	for i := 0; i < len(r); i++ {
		if unicode.IsDigit(r[i]) {
			count = count*10 + int(r[i]-'0') // consecutive digits are multiplied by 10
			continue
		}
		for count > 0 {
			sb.WriteRune(r[i])
			count--
		}
	}
	return sb.String()
}
