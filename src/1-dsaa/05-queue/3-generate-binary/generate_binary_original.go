package main

// Time complexity: O(n)
// Space complexity: O(n)
func generateBinaryOriginal(input int) []string {
	// not 0 on purpose as we initialize the queue with 1
	if input <= 1 {
		panic("not supported")
	}

	generatedNumbers := make([]string, input) // capacity not 0 on purpose
	queue := []string{"1"}
	for i := 0; i < input; i++ {
		generatedNumbers[i] = queue[0] // peak
		queue = queue[1:]              // dequeue

		// generate more
		queue = append(queue, generatedNumbers[i]+"0") // enqueue
		queue = append(queue, generatedNumbers[i]+"1")
	}

	return generatedNumbers
}
