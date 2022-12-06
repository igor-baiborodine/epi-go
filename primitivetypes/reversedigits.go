package primitivetypes

import (
	"math"
	"strconv"
)

// ReverseDigitsViaString returns the value of x with its digits in reversed order.
// The time/space complexity is O(n) where n is the number of digits in x.
func ReverseDigitsViaString(x int) int {
	if math.Abs(float64(x)) < 10 {
		return x
	}
	s := []rune(strconv.Itoa(x))
	i, j := 0, len(s)-1

	if x < 0 {
		i = 1
	}
	for i < j {
		s[i], s[j] = s[j], s[i]
		i, j = i+1, j-1
	}
	r, _ := strconv.Atoi(string(s))
	return r
}

// ReverseDigits returns the value of x with its digits in reversed order.
// The time is O(n) where n is the number of digits in x.
func ReverseDigits(x int) (r int) {
	r, xRemaining := 0, int(math.Abs(float64(x)))

	for xRemaining != 0 {
		r = r*10 + xRemaining%10
		xRemaining /= 10
	}
	if x < 0 {
		return -r
	}
	return r
}
