package main

import (
	"fmt"
)

// Time complexity: O(nmlogm)
// Space complexity: O(m)
func findTopKFrequentNumbers(input []int, k int) []int {
	frequencies := make(map[int]int, 0)
	for _, v := range input {
		frequencies[v]++
	}

	maxHeap := make([]int, 0)
	maxHeap = append(maxHeap, input...)
	buildMaxHeap(maxHeap, frequencies)

	outputDist := make(map[int]struct{}, 0)
	output := make([]int, 0)
	for len(maxHeap) > 0 && len(output) < k {
		v := maxHeap[0] // peak
		if _, ok := outputDist[v]; !ok {
			output = append(output, v)
			outputDist[v] = struct{}{}
		}
		maxHeap = maxHeap[1:] // pop
		buildMaxHeap(maxHeap, frequencies)
	}
	return output
}

func main() {
	input := []int{
		1, 3, 5, 12, 11, 12, 11,
	}
	k := 2
	fmt.Println(findTopKFrequentNumbers(input, k))
}

// Boilerplate coppied from the heap chapter
func maxHeapify(arr []int, i int, frequencies map[int]int) {
	smallest := i    // Initialize largest as root.
	left := 2*i + 1  // Left child index.
	right := 2*i + 2 // Right child index.

	// Check if the left child exists and is smaller than the root.
	if left < len(arr) && frequencies[arr[left]] > frequencies[arr[smallest]] {
		smallest = left
	}
	// Check if the right child exists and is smaller than the current smallest.
	if right < len(arr) && frequencies[arr[right]] > frequencies[arr[smallest]] {
		smallest = right
	}

	// If smallest is not root, swap it with the smallest and continue heapifying.
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		maxHeapify(arr, smallest, frequencies)
	}
}

func buildMaxHeap(arr []int, frequencies map[int]int) {
	// Start from the last non-leaf node and move upwards.
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, frequencies)
	}
}
