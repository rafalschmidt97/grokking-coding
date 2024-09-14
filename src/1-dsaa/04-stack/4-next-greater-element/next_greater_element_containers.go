package main

import "container/list"

func nextGreaterElementContainers(input []int) []int {
	result := make([]int, len(input)) // Initialize result array
	stack := list.New()               // Use a linked list as the stack

	// Iterate over each element in reverse order
	for i := len(input) - 1; i >= 0; i-- {
		// Pop elements from the stack until we find a greater element or stack is empty
		for stack.Len() > 0 && stack.Back().Value.(int) <= input[i] {
			stack.Remove(stack.Back()) // Pop the stack if the top is not greater
		}

		// If the stack is empty, no greater element was found, so set -1
		if stack.Len() == 0 {
			result[i] = -1
		} else {
			// The top of the stack is the next greater element
			result[i] = stack.Back().Value.(int)
		}

		// Push the current element onto the stack
		stack.PushBack(input[i])
	}

	return result
}
