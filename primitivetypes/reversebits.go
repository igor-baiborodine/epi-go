package primitivetypes

import "math/bits"

// ReverseBits returns the value of x with its bits in reversed order.
// The time complexity is O(n) where n is the binary word size.
// The space complexity is O(1).
func ReverseBits(x uint64) uint64 {
	for i := uint64(0); i < 64/2; i++ {
		x = SwapBits(x, i, 63-i)
	}
	return x
}

// ReverseBitsReverse64 returns the value of x with its bits in reversed order.
// The Reverse64() function from the math/bits package is used to reverse the binary word.
// The time complexity is O(1).
func ReverseBitsReverse64(x uint64) uint64 {
	return bits.Reverse64(x)
}
