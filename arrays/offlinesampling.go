package arrays

import (
	"math/rand"
	"time"
)

// RandomSampling performs random sampling of k elements from the input slice s.
// It modifies the original slice s in the process by swapping elements.
//
// Time Complexity: O(k), where k is the number of elements to sample. This is because the function
// performs k iterations, and each iteration involves a constant amount of work (swapping elements).
//
// Space Complexity: O(1), as the function operates in-place on the input slice and does not
// require any additional data structures.
func RandomSampling(k int, s []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < k; i++ {
		ri := i + r.Intn(len(s)-i)
		s[i], s[ri] = s[ri], s[i]
	}
	return s[:k]
}
