package primitivetypes

// ClosestInt returns integer closest to non-negative integer x with the same weight.
// The time complexity is O(n), where n is the integer width.
// The space complexity is O(1).
func ClosestInt(x uint64) (ci uint64, ok bool) {
	// x is assumed to be non-negative, so the leading bit is 0.
	// Therefore, limiting the range to 63 less significant bits.
	const numUnsignBits = 63

	for i := uint64(0); i < numUnsignBits; i++ {
		if (x>>i)&1 != x>>(i+1)&1 {
			x ^= 1<<i | 1<<(i+1)
			return x, true
		}
	}
	// All bits of x are either 0 or 1.
	return 0, false
}
