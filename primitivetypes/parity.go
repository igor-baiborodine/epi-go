package primitivetypes

// ParityBruteForce returns 1 if the number of 1s in x is odd, otherwise O.
// The time complexity is O(n) where n is the number of bits in x.
// The space complexity is O(1).
func ParityBruteForce(x uint64) (p uint16) {
	for x != 0 {
		p ^= uint16(x & 1)
		x >>= 1
	}
	return p
}
