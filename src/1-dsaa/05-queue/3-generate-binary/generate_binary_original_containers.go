package main

import "container/list"

// Time complexity: O(n)
// Space complexity: O(n)
func generateBinaryOriginalContainers(input int) []string {
	// If the input is less than 2, panic (as per the original logic)
	if input <= 1 {
		panic("not supported")
	}

	generatedNumbers := make([]string, input)
	queue := list.New() // Use list to simulate the queue
	queue.PushBack("1") // Initialize queue with "1"

	for i := 0; i < input; i++ {
		// Dequeue the front element
		current := queue.Front().Value.(string)
		queue.Remove(queue.Front())

		// Store the current binary number
		generatedNumbers[i] = current

		// Enqueue the next two numbers (current + "0", current + "1")
		queue.PushBack(current + "0")
		queue.PushBack(current + "1")
	}

	return generatedNumbers
}
