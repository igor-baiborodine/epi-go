package arrays

// DutchflagSubarrays segregates elements of the slice s into three groups:
// elements less than, equal to, and greater than the pivot s[p].
// It returns a new slice where all elements less than the pivot come first,
// followed by elements equal to the pivot, and finally elements greater than the pivot.
//
// Parameters:
// - s: Slice of integers to be partitioned.
// - p: Index of the pivot element.
//
// Returns:
// - A new reordered slice after partitioning.
//
// Time Complexity: O(n), where n is the length of the slice s.
// Space Complexity: O(n), due to the use of additional slices.
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

// DutchflagTwoPasses rearranges elements of the slice s such that all elements
// less than the pivot s[p] come first, followed by elements equal to the pivot,
// and finally elements greater than the pivot. It modifies and returns the n
// slice in-place using two passes.
//
// Parameters:
// - s: Slice of integers to be partitioned.
// - p: Index of the pivot element.
//
// Returns:
// - The same slice s after in-place partitioning.
//
// Time Complexity: O(n), where n is the length of the slice s.
// Space Complexity: O(1), as no additional space is used aside from variables.
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
