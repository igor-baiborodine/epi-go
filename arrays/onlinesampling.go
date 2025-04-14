package arrays

import (
	"math/rand"
	"time"

	"github.com/ipfs/boxo/routing/http/types/iter"
)

// OnlineRandomSampling generates a random sample of size k from a stream of integers
// represented by the given iterator. The sampling is performed online, meaning that
// the input stream does not need to be stored in memory.
//
// The algorithm uses reservoir sampling to ensure that each element in the input stream
// has an equal probability of being included in the sample.
//
// Time complexity: O(n), where n is the total number of elements in the input stream.
// The method iterates through the entire stream once, performing constant time
// operations per element.
//
// Space complexity: O(k), where k is the size of the sample. Only the first k elements
// and a few additional variables are stored in memory during processing.
func OnlineRandomSampling(k int, it *iter.SliceIter[int]) []int {
	runningSample := make([]int, k)
	// Stores the first k elements of the stream
	for i := 0; it.Next() && i < k; i++ {
		runningSample[i] = it.Val()
	}
	numSeenSoFar := k
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for it.Next() {
		curVal := it.Val()
		numSeenSoFar++
		// Generate a random number in [0, numSeenSoFar], and if this number is in
		// [0, k-1], we replace that element from the sample with curVal.
		idxToReplace := r.Intn(numSeenSoFar)
		if idxToReplace < k {
			runningSample[idxToReplace] = curVal
		}
	}
	return runningSample
}
