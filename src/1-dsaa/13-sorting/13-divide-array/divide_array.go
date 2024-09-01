package main

import (
	"fmt"
	"sort"
)

// Time complexity: O(nlogn)
//
// Space complexity: O(n), for current array
func divideArray(input []int, k int) [][]int {
	result := make([][]int, 0)

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	currentArray := make([]int, 0)
	for i, v := range input {
		currentArray = append(currentArray, v)
		if (i+1)%3 == 0 {
			if currentArray[2]-currentArray[0] <= k {
				result = append(result, currentArray)
			}
			currentArray = make([]int, 0) // clear
		}
	}

	return result
}

func main() {
	input := []int{
		2, 6, 4, 9, 3, 7, 3, 4, 1,
	}
	k := 3
	fmt.Println(divideArray(input, k))
}
