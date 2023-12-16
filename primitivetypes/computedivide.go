package primitivetypes

// ComputeDivide returns integer that is result of dividing two non-negative integers.
// The total time complexity of addition is 0(n), where n is the width of the operands.
func ComputeDivide(x, y uint64) uint64 {
	var quo uint64
	pow := 32
	yPow := y << pow

	for x >= y {
		for yPow > x {
			yPow = yPow >> 1
			pow = pow - 1
		}
		quo = quo + (1 << pow)
		x = x - yPow
	}
	return quo
}
