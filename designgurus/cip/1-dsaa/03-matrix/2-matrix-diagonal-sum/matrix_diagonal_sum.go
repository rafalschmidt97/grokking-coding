package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(1)
func matrixDiagonalSum(input [][]int) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		sum = sum + input[i][i] // add left diagonal

		isEvenMatrix := len(input)%2 == 0
		middlePointI := len(input) / 2 // starts from 0
		if isEvenMatrix || i != middlePointI {
			sum = sum + input[i][len(input)-1-i]
		}

	}

	return sum
}

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrixDiagonalSum(input))
}
