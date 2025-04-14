package arrays

func NextPermutation(p []int) []int {
	// NextPermutation generates the next lexicographical permutation of the slice `p`.
	// If the slice is already the largest permutation, it returns an empty slice.
	//
	// Algorithm:
	// 1. Traverse the array from right to left to find the first decreasing element (index `n`).
	//	This marks the pivot point where the next larger permutation can start. If no such index is found,
	//	it implies the array is the largest permutation, and an empty slice is returned.
	// 2. Starting from the end of the array, find the smallest element larger than `p[n]` (index `i`).
	//	This ensures the next permutation is the smallest possible increase from the current permutation.
	// 3. Swap `p[n]` and `p[i]` to move the next larger element into the pivot point.
	// 4. Reverse the subarray starting from `n+1` to reorder it into the smallest lexicographical order.
	//
	// Time Complexity:
	// O(n), where n is the length of the slice. The algorithm involves:
	// - One traversal to find the pivot (O(n)).
	// - One traversal to find the correct element to swap (O(n)).
	// - One traversal to reverse a portion of the array (O(n)).
	// Overall complexity is O(n).
	//
	// Space Complexity:
	// O(1), as the operations are performed in-place without requiring additional memory.
	k := len(p) - 2
	for k >= 0 && p[k] >= p[k+1] {
		k--
	}
	if k == -1 {
		return []int{} // p is the last permutation
	}
	// Swap the smallest entry after index n that is greater than p[n].
	// We exploit the fact that p[n+1:len(p)] is decreasing so if we search in reverse order,
	// the first entry that is p[n] is the smallest such entry.
	for i := len(p) - 1; i > k; i-- {
		if p[i] > p[k] {
			p[i], p[k] = p[k], p[i]
			break
		}
	}
	// Since p[n+1:len(p)) is in decreasing order, we can
	// build the smallest dictionary ordering of this subarray by reversing it.
	for i, j := k+1, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
	return p
}
