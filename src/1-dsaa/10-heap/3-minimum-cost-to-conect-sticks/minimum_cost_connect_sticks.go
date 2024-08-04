package main

import "fmt"

// Time complexity: O(nlogn)
// Space complexity: O(n)
func minimumCostToConnectSticks(input []int) int {
	minHeap := append([]int{}, input...)
	buildMinHeap(minHeap)

	costToConnect := 0
	for len(minHeap) >= 2 {
		v1 := minHeap[0]      // peak
		minHeap = minHeap[1:] // pop
		buildMinHeap(minHeap)
		v2 := minHeap[0] // peak
		combined := v1 + v2
		costToConnect += combined
		minHeap = minHeap[1:]               // pop
		minHeap = append(minHeap, combined) // push 1
		buildMinHeap(minHeap)
	}

	return costToConnect
}

func main() {
	input := []int{2, 4, 3}
	fmt.Println(minimumCostToConnectSticks(input))
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
