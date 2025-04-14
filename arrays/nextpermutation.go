package arrays

// NextPermutation generates the next permutation of the input slice p in lexicographic order.
// If the slice is already the last permutation, it returns an empty slice.
//
// Steps:
//  1. Identify the rightmost index k where p[k] < p[k+1]. This marks the position where the next larger permutation can start.
//     If no such index exists (k == -1), it means the array is already the largest permutation.
//  2. Find the smallest number larger than p[k] to the right of k. Since the subarray p[k+1:] is in decreasing order,
//     the first such number found in reverse traversal is optimal.
//  3. Swap p[k] with this number.
//  4. Reverse the subarray p[k+1:] to transform it into the smallest lexicographical order.
//
// Time Complexity:
// O(n), where n is the length of the array. The algorithm involves one pass to identify the pivot,
// another pass to find the element to swap, and finally one pass to reverse the subarray.
//
// Space Complexity:
// O(1), as the operations are performed in-place without requiring additional memory.
func NextPermutation(p []int) []int {
	k := len(p) - 2
	for k >= 0 && p[k] >= p[k+1] {
		k--
	}
	if k == -1 {
		return []int{} // p is the last permutation
	}
	// Swap the smallest entry after index k that is greater than p[k].
	// We exploit the fact that p[k+1:] is decreasing so if we search in reverse order,
	// the first entry that is p[k] is the smallest such entry.
	for i := len(p) - 1; i > k; i-- {
		if p[i] > p[k] {
			p[i], p[k] = p[k], p[i]
			break
		}
	}
	// Since p[k+1:] is in decreasing order, we can
	// build the smallest dictionary ordering of this subarray by reversing it.
	for i, j := k+1, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
	return p
}
