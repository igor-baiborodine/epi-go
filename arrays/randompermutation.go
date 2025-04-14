package arrays

// ComputeRandomPermutation generates a random permutation of the integers [0, n-1].
// It initializes a slice of size n, fills it with integers from 0 to n-1, and then
// performs a shuffling operation using OfflineRandomSampling.
//
// Time Complexity: O(n)
// Space Complexity: O(n) (for the slice to store the permutation)
func ComputeRandomPermutation(n int) []int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	return OfflineRandomSampling(n, perm)
}
