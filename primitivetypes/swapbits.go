package primitivetypes

// SwapBits returns the value of x with the bits swapped at indices i and j.
// The time and space complexity is O(1), independent of binary word size.
func SwapBits(x, i, j uint64) uint64 {
	// Extract the i-tf and j-th bits, and see if they differ.
	if (x>>i)&1 != (x>>j)&1 {
		// i-th and j-th bits differ. We will swap them by flipping their values.
		// Select the bits to flip with bitMask. Since x^1 = <0 when x = 1 and 1
		// when x = 0, we can perform the flip XOR .
		x ^= 1<<i | 1<<j
	}
	return x
}
