package strings

// ReplaceRemove replaces each "a" with two "d"s and removes each "b"
// from a slice of characters where n is the number of characters
// that the operation is to be applied to.
// The input slice can only contain ASCII characters.
// The time complexity is O(n) and space complexity is O(1).
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
