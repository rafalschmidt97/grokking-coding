package main

import "fmt"

// Traverse through each element, cumulatively summing up
// the values as we proceed and placing the running total at
// the corresponding index in the resulting array.
//
// Time complexity: O(n)
// Space complexity: O(1), only one counter
func runningSum(input []int) {
	if len(input) <= 1 {
		return
	}

	sum := input[0]
	for i := 1; i < len(input); i++ {
		sum = sum + input[i]
		input[i] = sum
	}
}

func main() {
	input := []int{2, 3, 5, 1, 6}
	// var input []int // array len works with nils
	runningSum(input)
	fmt.Println(input)
}
