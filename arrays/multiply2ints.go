package arrays

import (
	"math"
)

// Multiply2Ints multiplies two slices of digits representing signed integers.
// Each slice represents an integer with the most significant digit at the beginning.
// The result is a slice of digits representing the product of the two integers.
//
// Parameters:
// - n1: A slice of int8 digits representing the first integer.
// - n2: A slice of int8 digits representing the second integer.
//
// Returns:
// - A slice of int8 digits representing the product of the two input integers.
//
// Notes:
// - The function handles sign multiplication and ensures the correct sign of the result.
// - Excess leading zeros in the result are removed before returning the final slice.
//
// Time Complexity: O(n1 * n2), where n1 and n2 are the lengths of the input slices sn1 and sn2,
// as each digit of sn1 is multiplied with each digit of sn2.
// Space Complexity: O(n1 + n2), since the result can have at most n1 + n2 digits in the worst case.
func Multiply2Ints(sn1, sn2 []int8) []int8 {
	if len(sn1) == 1 && sn1[0] == 0 || len(sn2) == 1 && sn2[0] == 0 {
		return []int8{0}
	}

	sign := combinedSign(int(sn1[0]), int(sn2[0]))
	sn1[0] = int8(math.Abs(float64(sn1[0])))
	sn2[0] = int8(math.Abs(float64(sn2[0])))
	s := make([]int8, len(sn1)+len(sn2))

	for i := len(sn1) - 1; i >= 0; i-- {
		for j := len(sn2) - 1; j >= 0; j-- {
			s[i+j+1] += sn1[i] * sn2[j]
			s[i+j] += s[i+j+1] / 10
			s[i+j+1] %= 10
		}
	}

	firstNotZero := 0
	for i, n := range s {
		if n != 0 {
			firstNotZero = i
			break
		}
	}
	s = s[firstNotZero:]
	if len(s) == 0 {
		return []int8{0}
	}
	s[0] = s[0] * sign
	return s
}

func combinedSign(a, b int) int8 {
	if a == 0 || b == 0 {
		return 0
	}
	if (a ^ b) < 0 {
		return -1
	}
	return 1
}
