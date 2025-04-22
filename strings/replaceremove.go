package strings

// ReplaceRemove modifies an input slice of strings s based on the following rules:
// 1. Replaces every occurrence of "a" with two "d"s.
// 2. Removes every occurrence of "b".
// The function operates in two passes:
// - Forward pass: Counts the occurrences of "a" and removes all "b"s in-place.
// - Backward pass: Modifies the slice to replace "a" with "dd", starting from the end.
//
// Parameters:
// - n (int): The number of elements to consider in s.
// - s ([]string): The slice of strings to modify.
//
// Returns:
// - []string: The modified slice containing the changes specified.
//
// Time Complexity:
// - O(n): The function processes the input slice twice - forward pass and backward pass.
//
// Space Complexity:
//   - O(1): The function operates in-place without requiring additional data structures,
//     apart from a small constant amount of extra variables.
func ReplaceRemove(n int, s []string) []string {
	var wi, ac int
	//fmt.Println("[1] n, s:", n, s)

	// Forward pass: count "a"s and remove "b"s
	for i := 0; i < n; i++ {
		if s[i] == "a" {
			ac++
		}
		if s[i] != "b" {
			if wi < i {
				s[wi] = s[i] // overwrite previous "b"
			}
			wi++
		}
		//fmt.Println("[2] wi, i, ac, s:", wi, i, ac, s)
	}

	// Backward pass: replace "a"s with two "d"s moving from the end
	ci := wi - 1
	wi = wi + ac - 1
	fs := wi + 1
	//fmt.Println("[3] ci, wi, fs:", ci, wi, fs)

	for ci >= 0 {
		if s[ci] == "a" {
			s[wi-1] = "d"
			s[wi] = "d"
			wi = wi - 2
		} else {
			s[wi] = s[ci]
			wi--
		}
		ci--
		//fmt.Println("[4] wi, ci, s:", wi, ci, s)
	}
	return s[:fs]
}
