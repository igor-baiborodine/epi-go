package arrays

// PlusOne takes a slice of digits that encode a non-negative decimal integer D
// and updates it to represent the integer D + 1.
//
// Parameters:
// - s: A slice of uint8, where each element is a digit (0-9) of the number.
//
// Returns:
// - A slice of uint8 representing the updated number.
//
// Time Complexity: O(n), where n is the length of the n slice. Each digit is processed once.
// Space Complexity: O(1), as the function operates in-place and does not use significant additional memory.
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
