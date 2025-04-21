package arrays

import "math"

// IsValidSudoku checks if a given Sudoku board is valid.
// A Sudoku board is considered valid if:
// - Each row contains no duplicate numbers (other than 0, which represents empty cells).
// - Each column contains no duplicate numbers (other than 0).
// - Each sub-grid (region) contains no duplicate numbers (other than 0).
//
// Arguments:
//   - board: A 2D slice where each element represents a cell in a Sudoku grid.
//
// Returns:
//   - true if the Sudoku board is valid, false otherwise.
//
// Time Complexity: O(n^2), where n is the dimension of the board (e.g., 9 for a 9x9 board).
//   - Checking rows and columns takes O(n^2) (O(n^2) for visiting every cell once across rows and columns).
//   - Checking regions also takes O(n^2) for iterating through each cell in the sub-grids.
//
// Space Complexity: O(n):
//   - The auxiliary space for the isPresent slice used in hasDuplicates, which scales with the number of rows/columns.
func IsValidSudoku(board [][]int) bool {
	// Check row constraints.
	for i := 0; i < len(board); i++ {
		if hasDuplicates(board, i, i+1, 0, len(board)) {
			return false
		}
	}

	// Check column constraints.
	for j := 0; j < len(board); j++ {
		if hasDuplicates(board, 0, len(board), j, j+1) {
			return false
		}
	}

	// Check region constraints.
	regionSize := int(math.Sqrt(float64(len(board))))
	for i := 0; i < regionSize; i++ {
		for j := 0; j < regionSize; j++ {
			if hasDuplicates(board, i*regionSize, (i+1)*regionSize-1, j*regionSize, (j+1)*regionSize-1) {
				return false
			}
		}
	}
	return true
}

func hasDuplicates(board [][]int, startRow, endRow, startCol, endCol int) bool {
	isPresent := make([]bool, len(board)+1)

	for i := startRow; i < endRow; i++ {
		for j := startCol; j < endCol; j++ {
			if board[i][j] != 0 && isPresent[board[i][j]] {
				return true
			}
			isPresent[board[i][j]] = true
		}
	}
	return false
}
