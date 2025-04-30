package strings

// ReverseWords reverses the order of words in a string while preserving the order of characters
// within each word. Words are separated by single spaces. The function first reverses the entire
// string, then reverses each word individually.
//
// Parameters:
// - s: The input string containing words separated by spaces.
//
// Returns:
// - A string with all words in reverse order but characters within each word preserved.
//
// Time Complexity: O(n) where n is the length of the input string
// Space Complexity: O(n) due to the conversion of string to rune slice
func ReverseWords(s string) string {
	r := []rune(s)
	reverse(r, 0, len(s)-1)

	start := 0
	for i := 0; i < len(r); i++ {
		if r[i] == ' ' {
			reverse(r, start, i-1)
			start = i + 1
		}
	}
	reverse(r, start, len(r)-1)
	return string(r)
}

func reverse(r []rune, start, end int) {
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
}
