package main

import "fmt"

// A simple algorithm is to run two loops: the outer loop picks all elements one by one,
// and the inner loop looks for the first greater element for the element picked by the outer loop.
// However, this algorithm has a time complexity of O(n^2).
//
// We can use a more optimized approach using Stack data structure.
// The algorithm will leverage the nature of the stack data structure,
// where the most recently added (pushed) elements are the first ones to be
// removed (popped). Starting from the end of the array, the algorithm always
// maintains elements in the stack that are larger than the current element.
// This way, it ensures that it has a candidate for the "next larger element".
// If there is no larger element, it assigns -1 to that position.
// It handles each element of the array only once, making it an efficient solution.
//
// Time complexity: O(n) as every element is pushed and popped from the stack exactly once.
// Space complexity: O(n)
func nextGreaterElemenet(input []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(input))

	for i := len(input) - 1; i >= 0; i-- {
		foundGreater := false

		for !foundGreater && len(stack) != 0 {
			top := stack[len(stack)-1]

			if top > input[i] {
				// consider setting in outer loop and here only pop if necessary
				result[i] = top
				stack = append(stack, input[i])
				foundGreater = true
			} else {
				stack = stack[:len(stack)-1] // pop
			}
		}

		if !foundGreater {
			result[i] = -1
			stack = append(stack, input[i])
		}
	}
	return result
}

func main() {
	input := []int{4, 5, 2, 25}
	fmt.Println(nextGreaterElemenet(input))
}
