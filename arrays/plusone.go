package arrays

// PlusOne updates a slice of digits that encodes
// a non-negative decimal integer D to represent the integer D + 1.
// The time complexity is O(n) and space complexity is O(1).
func PlusOne(s []uint8) []uint8 {
	s[len(s)-1]++

	for i := len(s) - 1; i >= 1; i-- {
		if s[i] != 10 {
			break
		}
		s[i] = 0
		s[i-1]++
	}
	if s[0] == 10 {
		s[0] = 1
		s = append(s, 0)
	}
	return s
}
