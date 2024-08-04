package main

import "fmt"

// Time complexity: O(N+M), worst case O(2N) = O(N)
// Space complexity: O(N+M), worst case O(2N) = O(N)
func uniqueNumbers(input []int) bool {
	occurancesSet := make(map[int]int, 0)
	uniqueSet := make(map[int]struct{}, 0)

	for _, number := range input {
		occurancesSet[number]++
	}

	for _, occurances := range occurancesSet {
		if _, ok := uniqueSet[occurances]; ok {
			return false
		}

		uniqueSet[occurances] = struct{}{}
	}

	return true
}

func main() {
	input := []int{4, 5, 4, 6, 6, 6}
	fmt.Println(uniqueNumbers(input))
}
