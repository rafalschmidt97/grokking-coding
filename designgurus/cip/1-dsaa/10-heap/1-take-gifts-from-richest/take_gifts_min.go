package main

import (
	"math"
)

// Min heap implementation
// Time complexity: O(nlogn)
// Space complexity: O(n)
func takeGiftsFromRichestMin(input []int, k int) int {
	// We're using a min heap, but by inserting negative numbers,
	// we effectively convert it into a max heap.
	minHeap := make([]int, 0)

	// Populate the max heap with the negative of values and build.
	for _, v := range input {
		minHeap = append(minHeap, -v)
	}
	buildMinHeap(minHeap)

	for i := 0; i < k; i++ {
		max := -minHeap[0] // pop first (smallest) and make it biggest
		maxSqrt := int(math.Floor(math.Sqrt(float64(max))))
		minHeap = append(minHeap[1:], -maxSqrt) // push with shrink
		minHeapify(minHeap, 0)
	}

	remaining := 0
	for _, v := range minHeap {
		remaining += -v
	}
	return remaining
}

// Boilerplate

func minHeapify(arr []int, i int) {
	smallest := i    // Initialize largest as root.
	left := 2*i + 1  // Left child index.
	right := 2*i + 2 // Right child index.

	// Check if the left child exists and is smaller than the root.
	if left < len(arr) && arr[left] < arr[smallest] {
		smallest = left
	}
	// Check if the right child exists and is smaller than the current smallest.
	if right < len(arr) && arr[right] < arr[smallest] {
		smallest = right
	}

	// If smallest is not root, swap it with the smallest and continue heapifying.
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		minHeapify(arr, smallest)
	}
}

func buildMinHeap(arr []int) {
	// Start from the last non-leaf node and move upwards.
	for i := len(arr)/2 - 1; i >= 0; i-- {
		minHeapify(arr, i)
	}
}
