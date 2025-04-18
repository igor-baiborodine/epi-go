package arrays

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// GenerateNonuniformRandomNumber selects a random number from the input array v
// based on the provided probabilities p. Each probability in p corresponds to
// the likelihood of selecting the number at the same index in v.
//
// Parameters:
// - v: A slice of integers representing the values to choose from.
// - p: A slice of probabilities (in the range [0, 1]) corresponding to each value in v.
//
// Returns:
// - An integer from the slice v, chosen based on the probabilities in p.
//
// Notes:
//   - The sums of the probabilities in p must equal 1.
//   - The function generates a uniformly distributed random value in [0, 1) and determines which
//     interval (derived from the prefix sums of p) this value falls into, to select the corresponding element in v.
//
// Time Complexity: O(n log n)
// - O(n) to compute the prefix sums of the probabilities.
// - O(log n) for the binary search to determine the appropriate interval.
//
// Space Complexity: O(n)
// - Space is used to store the prefix sums of the probabilities.
func GenerateNonuniformRandomNumber(v []int, p []float64) int {
	prefixSumOfProbabilities := []float64{0.0}
	// Creating the endpoints for the intervals corresponding to the probabilities.
	for _, probability := range p {
		prefixSumOfProbabilities =
			append(prefixSumOfProbabilities, prefixSumOfProbabilities[len(prefixSumOfProbabilities)-1]+probability)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	uniform01 := r.Float64()
	fmt.Println("uniform01:", uniform01)

	// Find the index of the interval that uniform01 lies in.
	it := sort.Search(len(prefixSumOfProbabilities), func(i int) bool {
		return prefixSumOfProbabilities[i] >= uniform01
	})

	if it < 0 {
		// We want the index of the first element in the array which is
		// greater than the key.
		//
		// When a key is not present in the array, binary search
		// returns the negative of 1 plus the smallest index whose entry
		// is greater than the key.
		//
		// Therefore, if the return value is negative, by taking its absolute
		// value and adding 1 to it, we get the desired index.
		intervalIdx := int((math.Abs(float64(it)) - 1) - 1)
		return v[intervalIdx]
	} else {
		// We have it >= 0, i.e., uniform01 equals an entry in prefixSumOfProbabilities.
		//
		// Because we uniform01 is a random float64, the probability of it
		// equalling an endpoint in prefixSumOfProbabilities is exceedingly low.
		// However, it is not 0, so to be robust we must consider this case.
		return v[it-1]
	}
}
