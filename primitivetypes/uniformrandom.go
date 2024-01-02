package primitivetypes

import "math/rand"

// UniformRandom returns, as an uint64, a non-negative uniform random number between x and y, inclusive.
func UniformRandom(min, max uint64) uint64 {
	numOfOutcomes := max - min + 1
	res := numOfOutcomes

	for res >= numOfOutcomes {
		res = 0
		for i := 0; 1<<i < numOfOutcomes; i++ {
			res = res<<1 | uint64(rand.Intn(2))
		}
	}
	return res + min
}
