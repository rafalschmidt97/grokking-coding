package main

import "container/list"

func sortStackContainers(input []int) []int {
	inputStack := list.New()  // The input stack
	sortedStack := list.New() // The auxiliary stack for sorted elements

	// Push all input elements onto the inputStack
	for _, v := range input {
		inputStack.PushBack(v)
	}

	// Sort the stack using the auxiliary stack (sortedStack)
	for inputStack.Len() > 0 {
		// Pop the top element from inputStack
		current := inputStack.Back().Value.(int)
		inputStack.Remove(inputStack.Back())

		// Move elements from sortedStack to inputStack if they are larger than the current element
		for sortedStack.Len() > 0 && sortedStack.Back().Value.(int) > current {
			inputStack.PushBack(sortedStack.Back().Value.(int))
			sortedStack.Remove(sortedStack.Back())
		}

		// Push the current element onto sortedStack
		sortedStack.PushBack(current)
	}

	// Collect elements from sortedStack in the correct (ascending) order to return
	result := make([]int, 0, sortedStack.Len())
	for sortedStack.Len() > 0 {
		result = append(result, sortedStack.Front().Value.(int))
		sortedStack.Remove(sortedStack.Front()) // Collect from the front (smallest first)
	}

	return result
}
