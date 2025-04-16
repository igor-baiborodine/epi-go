package arrays

import (
	"math/rand"
	"time"
)

// RandomSubset generates a random subset of size k from the range [0, n-1].
// It uses a map to keep track of changed elements, effectively creating a random permutation
// of the first k elements in the range.
//
// Time Complexity: O(k)
// The loop runs k iterations, with each iteration performing constant time operations.
//
// Space Complexity: O(k)
// The space used is proportional to the number of entries in the map, which is at most k.
func RandomSubset(n int, k int) []int {
	changedElements := map[int]int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < k; i++ {
		// Generate random int in [i, n-1]
		randomIdx := i + r.Intn(n-i)
		elem1, ok1 := changedElements[randomIdx]
		elem2, ok2 := changedElements[i]

		if !ok1 && !ok2 {
			changedElements[randomIdx] = i
			changedElements[i] = randomIdx
		} else if !ok1 && ok2 {
			changedElements[randomIdx] = elem2
			changedElements[i] = randomIdx
		} else if ok1 && !ok2 {
			changedElements[i] = elem1
			changedElements[randomIdx] = i
		} else {
			changedElements[i] = elem1
			changedElements[randomIdx] = elem2
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = changedElements[i]
	}
	return result
}
