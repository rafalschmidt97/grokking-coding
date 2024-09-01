package main

import (
	"fmt"
	"sort"
)

// Time complexity: O(nlogn)
// Space complexity: O(n)
func sortByFrequency(input []int) {
	frequencies := make(map[int]int, 0)
	for _, v := range input {
		frequencies[v]++
	}
	sort.Slice(input, func(i, j int) bool {
		return frequencies[input[i]] <= frequencies[input[j]]
	})
}

func main() {
	input := []int{
		4, 4, 6, 2, 2, 2,
	}
	sortByFrequency(input)
	fmt.Println(input)
}
