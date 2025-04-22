package arrays

// GeneratePascalTriangle generates Pascal's Triangle up to the nth level.
// n specifies the number of rows in the triangle (0-indexed).
// Each row in the returned 2D slice represents a row in Pascal's Triangle.
//
// Time Complexity: O(n^2)
// - Two nested loops iterate to build the triangle:
//   - Outer loop runs n times (for each row).
//   - Inner loop runs up to i times for each row (summing up to n(n+1)/2 iterations).
//
// Space Complexity: O(n^2)
//   - The 2D slice pascalTriangle stores the entire triangle
//     with approximately (n(n+1))/2 elements in total.
func GeneratePascalTriangle(n int) [][]int {
	pascalTriangle := [][]int{}
	for i := 0; i < n; i++ {
		currentRow := []int{}
		for j := 0; j <= i; j++ {
			if 0 < j && j < i {
				currentRow = append(currentRow, pascalTriangle[i-1][j-1]+pascalTriangle[i-1][j])
			} else {
				currentRow = append(currentRow, 1)
			}
		}
		pascalTriangle = append(pascalTriangle, currentRow)
	}
	return pascalTriangle
}
