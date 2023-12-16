package primitivetypes

import "math/bits"

// ComputeMultiply returns integer that is result of multiplying two non-negative integers.
// The total time complexity of addition is O(n), where n is the width of the operands.
func ComputeMultiply(x, y uint64) uint64 {
	var sum uint64

	for x != 0 {
		// Examines each bit of x.
		if (x & 1) != 0 {
			sum, _ = bits.Add64(sum, y, 0)
		}
		x = x >> 1
		y = y << 1
	}
	return sum
}
