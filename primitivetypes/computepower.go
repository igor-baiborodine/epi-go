package primitivetypes

// ComputePower returns floating-point that is result of floating-point to the power of non-negative integer.
// The time complexity of addition is O(n).
func ComputePower(x float64, y int64) float64 {
	var res = 1.0
	pow := y

	if y < 0 {
		pow = -pow
		x = 1.0 / x
	}
	for pow != 0 {
		if pow&1 != 0 {
			res = res * x
		}
		x = x * x
		pow = pow >> 1
	}
	return res
}
