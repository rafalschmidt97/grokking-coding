package main

import "fmt"

// Time complexity: TODO:
// Space complexity: TODO:
func maxWithOnes(input [][]int) []int {
	maxIndex := -1
	maxCount := 0

	for i, v := range input { // rows
		rowCount := 0

		for _, v := range v { // rows
			if v == 1 {
				rowCount = rowCount + 1
			}
		}

		if rowCount > maxCount {
			maxIndex = i
			maxCount = rowCount
		}
	}

	return []int{maxIndex, maxCount}
}

func main() {
	input := [][]int{
		{1, 0},
		{1, 1},
		{0, 1},
	}
	fmt.Println(maxWithOnes(input))
}
