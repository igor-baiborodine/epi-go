package arrays

// RotateMatrix rotates a given square matrix 90 degrees clockwise in place.
//
// It performs a 4-way exchange to rotate the matrix layer by layer, starting
// from the outermost layer and moving inward. The matrix is modified in place.
//
// Time Complexity: O(n^2), where n is the size of the matrix (length of one side).
// This is because we iterate over almost all elements in the matrix once in nested loops.
//
// Space Complexity: O(1), as we do not use any additional space apart from the given matrix.
func RotateMatrix(m *[][]int) {
	matrixSize := len(*m) - 1
	for i := 0; i < len(*m)/2; i++ {
		for j := i; j < matrixSize-i; j++ {
			// Perform a 4-way exchange.
			temp := (*m)[i][j]
			(*m)[i][j] = (*m)[matrixSize-j][i]
			(*m)[matrixSize-j][i] = (*m)[matrixSize-i][matrixSize-j]
			(*m)[matrixSize-i][matrixSize-j] = (*m)[j][matrixSize-i]
			(*m)[j][matrixSize-i] = temp
		}
	}
}
