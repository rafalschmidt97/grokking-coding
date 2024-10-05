package main

import (
	"fmt"
	"sort"
)

// Time complexity: O(nlogn)
// Space Complexity: O(1), if pdqsort uses insertion sort, otherwise O(n)
func maxLenPairChain(input [][]int) int {
	sort.Slice(input, func(i, j int) bool {
		return input[i][1] < input[j][1]
	})

	counter := 1
	last := input[0]
	for i := 1; i < len(input); i++ {
		if last[1] < input[i][0] {
			last = input[i]
			counter++
		}
	}

	return counter
}

func main() {
	input := [][]int{{1, 2}, {3, 4}, {2, 3}}
	fmt.Println(maxLenPairChain(input))
}
