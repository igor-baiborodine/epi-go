package strings

import "strings"

// SnakeString converts a string into a "snake string" by rearranging characters in a snake-like pattern:
// - first line: characters at positions 1, 5, 9, ...
// - second line: characters at positions 0, 2, 4, ...
// - third line: characters at positions 3, 7, 11, ...
//
// Parameters:
//   - s: input string to be converted
//
// Returns:
//   - string: the snake-string representation of the input
//
// Time Complexity: O(n) where n is the length of the input string
// Space Complexity: O(n) to store the result string builder
func SnakeString(s string) string {
	r := []rune(s)
	sb := strings.Builder{}
	for i := 1; i < len(r); i = i + 4 {
		sb.WriteRune(r[i])
	}
	for i := 0; i < len(r); i = i + 2 {
		sb.WriteRune(r[i])
	}
	for i := 3; i < len(r); i = i + 4 {
		sb.WriteRune(r[i])
	}
	return sb.String()
}
