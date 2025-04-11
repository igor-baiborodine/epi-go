package arrays

// ApplyPermutation rearranges the elements of the slice s according to the permutation array p.
// Each index i in p indicates the new position of the element currently at index i in s.
//
// Parameters:
//   - p: A slice of integers representing the permutation.
//     p[i]contains the new position of the element at index i.
//   - s: A slice of integers (or values of any type) to be rearranged according to the permutation.
//
// The function modifies the input slice s in-place and modifies p during the process, but it restores
// p to its original state by the end of the function.
//
// Time Complexity: O(n), where n is the size of the slice s, as each element is visited at most once.
// Space Complexity: O(1), as the function operates in-place without using additional data structures.
func ApplyPermutation(p []int, s []int) {
	for i := range s {
		// Check if the element at index i has not been moved by checking if
		// p[i] is non-negative.
		next := i
		for p[next] >= 0 {
			temps := s[i]
			s[i] = s[p[next]]
			s[p[next]] = temps

			tempp := p[next]
			// Subtracts length of p from an entry in p to make it negative,
			// which indicates the corresponding move has been performed.
			p[next] = p[next] - len(p)
			next = tempp
		}
	}
	// Restore values in p.
	for i := range p {
		p[i] = p[i] + len(p)
	}
}
