package main

import (
	"fmt"
	"sort"
)

// Time complexity: O(nlogn)
//
// Space complexity: O(1), if sort implementation use inplace
// replace (not merge/quick sort). For instance heap/insertion.
func makeArrayEqual(input []int) int {
	// Sort the array in descending order
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	operations := 0

	// Iterate through the array and count the operations needed
	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] {
			operations += i
		}
	}

	return operations
}

func main() {
	input := []int{3, 5, 5, 2}
	fmt.Println(makeArrayEqual(input))
}
