package main

import (
	"fmt"
	"sort"
)

// Time complexity: O(mlogn), as sorting dominates
// Space complexity: O(n)
func distributeApples(input []int, capacity []int) int {
	sort.Slice(capacity, func(i, j int) bool { // sort desc
		return capacity[i] > capacity[j]
	})

	apples := 0
	for _, v := range input {
		apples += v
	}

	used := 0
	for _, v := range capacity {
		used++

		apples -= v

		if apples <= 0 {
			break
		}
	}

	return used
}

func main() {
	input := []int{
		2, 3, 1,
	}
	capacity := []int{
		4, 2, 5, 1,
	}
	fmt.Println(distributeApples(input, capacity))
}
