package main

import (
	"fmt"
)

// Time complexity: O(N)
// Space complexity: O(N)
func largestUniqueNumber(input []int) int {
	numberMap := make(map[int]int, 0)

	for _, number := range input {
		numberMap[number]++
	}

	highestValue := -1
	for number, occurances := range numberMap {
		if occurances == 1 {
			if highestValue < number {
				highestValue = number
			}
		}
	}

	return highestValue
}

func main() {
	input := []int{1, 2, 3}
	fmt.Println(largestUniqueNumber(input))
}
