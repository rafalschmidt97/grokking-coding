package main

// Time complexity: O(m + n)
// Space complexity: O(1)
func search2dMatrixOriginal(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := 0
	col := len(matrix[0]) - 1

	// Start from top right corner and move towards the target
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			// Target found
			return true
		} else if matrix[row][col] < target {
			// Move down
			row++
		} else {
			// Move left
			col--
		}
	}

	return false
}
