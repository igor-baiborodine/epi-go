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

// ParityBruteForceImproved returns 1 if the number of 1s in x is odd, otherwise O.
// The time complexity is O(k) where k is the number of bits set to 1 in x.
// The space complexity is O(1).
func ParityBruteForceImproved(x uint64) (p uint16) {
	for x != 0 {
		p ^= 1
		x &= x - 1
	}
	return p
}

// Parity returns 1 if the number of 1s in x is odd, otherwise O.
// The time complexity is O(log(n)) where n is the size of x.
// The space complexity is O(1).
func Parity(x uint64) (p uint16) {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return uint16(x & 1)
}
