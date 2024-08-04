package main

import "fmt"

// Time complexity: O(N)
// Space complexity: O(N)
func counteringElements(input []int) int {
	numbersSet := make(map[int]struct{}, 0)
	for _, number := range input {
		numbersSet[number] = struct{}{}
	}
	counter := 0
	for _, number := range input {
		if _, ok := numbersSet[number+1]; ok {
			counter++
		}
	}
	return counter
}

func main() {
	input := []int{4, 3, 1, 5, 6}
	fmt.Println(counteringElements(input))
}
