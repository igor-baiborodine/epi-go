package primitivetypes

import "math"

// PalindromeNumber determines if that integer's representation as a decimal string is a palindrome.
// The time complexity is O(n), and the space complexity is O(1).
func PalindromeNumber(x int) bool {
	if x <= 0 {
		return x == 0
	}
	numDigits := int(math.Floor(math.Log10(float64(x))) + 1)
	msdMask := int(math.Pow10(numDigits - 1))

	for i := 0; i < numDigits/2; i++ {
		if x/msdMask != x%10 {
			return false
		}
		x = x % msdMask // Remove the most significant digit of x.
		x = x / 10      // Remove the least significant digit of x.
		msdMask = msdMask / 100
	}
	return true
}
