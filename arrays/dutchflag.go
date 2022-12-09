package arrays

// DutchflagSubarrays reorders slice so that all elements
// less than s[i] appear first, followed by elements equal to s[i],
// followed by elements greater than s[i].
// The time complexity is O(n) and space complexity is O(n).
func DutchflagSubarrays(s []int, p int) []int {
	var r, w, b []int

	for _, v := range s {
		if v < s[p] {
			r = append(r, v)
		} else if v == s[p] {
			w = append(w, v)
		} else {
			b = append(b, v)
		}
	}
	s = append(r, w...)
	return append(s, b...)
}

// DutchflagTwoPasses reorders slice so that all elements
// less than s[i] appear first, followed by elements equal to s[i],
// followed by elements greater than s[i].
// The time complexity is O(n) and space complexity is O(1).
func DutchflagTwoPasses(s []int, p int) []int {
	pivot, smaller := s[p], 0
	// First pass: group elements smaller than pivot
	for i := range s {
		if s[i] < pivot {
			s[i], s[smaller] = s[smaller], s[i]
			smaller++
		}
	}
	// Second pass: group elements larger than pivot
	larger := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < pivot {
			break
		} else if s[i] > pivot {
			s[i], s[larger] = s[larger], s[i]
			larger--
		}
	}
	return s
}
