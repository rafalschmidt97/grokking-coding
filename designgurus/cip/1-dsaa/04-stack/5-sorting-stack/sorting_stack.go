package main

import "fmt"

// How does it work
// Input Stack: [34, 3, 31, 98, 92], tmpStack: [23]
// Input Stack: [34, 3, 31, 98], tmpStack: [23, 92]
// Input Stack: [34, 3, 31], tmpStack: [23, 92, 98]
// Input Stack: [34, 3, 98, 92], tmpStack: [23, 31]
// nput Stack: [34, 3], tmpStack: [23, 31, 92, 98]
// Input Stack: [34, 98, 92, 31, 23], tmpStack: [3]
// Input Stack: [34], tmpStack: [3, 23, 31, 92, 98]
// Input Stack: [98, 92], tmpStack: [3, 23, 31, 34]
// Sorted: [3, 23, 31, 34, 92, 98]
//
// Time complexity: O(n^2) in the worst case
// Space complexity: O(n)
func sortStack(input []int) []int {
	sortedInput := make([]int, 0, len(input))

	for len(input) > 0 {
		inputPeek := input[len(input)-1]
		input = input[:len(input)-1] // pop

		for len(sortedInput) > 0 {
			sortedInputPeek := sortedInput[len(sortedInput)-1]

			if inputPeek > sortedInputPeek {
				break
			}

			sortedInput = sortedInput[:len(sortedInput)-1] // pop
			input = append(input, sortedInputPeek)         // push
		}

		sortedInput = append(sortedInput, inputPeek) // push
	}

	return sortedInput
}

func main() {
	input := []int{34, 3, 31, 98, 92, 23}
	fmt.Println(sortStack(input))
}
