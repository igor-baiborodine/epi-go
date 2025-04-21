package arrays

import "math"

// MatrixInSpiralOrder takes a 2D matrix m and returns a slice of integers containing
// the elements of the matrix traversed in a spiral order, starting from the top-left corner.
// It iterates through layers of the matrix and processes each layer in a clockwise manner
// by calling the helper function matrixLayerInClockwise.
//
// Time Complexity: O(n^2), where n is the dimension of the matrix. This is because all elements
// of the matrix are traversed exactly once in a spiral order.
// Space Complexity: O(1) additional space, as the function operates directly on the given matrix
// and uses a slice to store the result (excluding the output slice).
func MatrixInSpiralOrder(m [][]int) []int {
	so := []int{}
	for offset := 0; offset < int(math.Ceil(float64(len(m))/2)); offset++ {
		matrixLayerInClockwise(m, offset, &so)
	}
	return so
}

func matrixLayerInClockwise(m [][]int, offset int, so *[]int) {
	// matrix m has odd dimension, and we are at its center.
	if offset == len(m)-offset-1 {
		*so = append(*so, m[offset][offset])
		return
	}
	for j := offset; j < len(m)-offset-1; j++ {
		*so = append(*so, m[offset][j])
	}
	for i := offset; i < len(m)-offset-1; i++ {
		*so = append(*so, m[i][len(m)-offset-1])
	}
	for j := len(m) - offset - 1; j > offset; j-- {
		*so = append(*so, m[len(m)-offset-1][j])
	}
	for i := len(m) - offset - 1; i > offset; i-- {
		*so = append(*so, m[i][offset])
	}
}
