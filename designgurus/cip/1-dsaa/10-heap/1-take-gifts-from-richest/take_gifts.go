package main

import (
	"fmt"
	"math"
)

// Max heap implementation
// Time complexity: O(nlogn)
// Space complexity: O(n)
func takeGiftsFromRichestMax(input []int, k int) int {
	maxHeap := append([]int{}, input...)
	buildMaxHeap(maxHeap)

	for i := 0; i < k; i++ {
		max := maxHeap[0] // pop first (largest)
		maxSqrt := int(math.Floor(math.Sqrt(float64(max))))
		maxHeap = append(maxHeap[1:], maxSqrt) // push with shrink
		// maxHeap[0] = maxHeap[len(maxHeap)-1] // Replace the root with the last element and shrink the heap.
		// maxHeap = maxHeap[:len(maxHeap)-1]
		// maxHeap = append(maxHeap, maxSqrt) // Add the square root value to the heap.
		maxHeapify(maxHeap, 0)
	}

	remaining := 0
	for _, v := range maxHeap {
		remaining += v
	}
	return remaining
}

// Boilerplate for singly linked list

// maxHeapify ensures the max heap property for the subtree rooted at index i.
func maxHeapify(arr []int, i int) {
	largest := i     // Initialize largest as root.
	left := 2*i + 1  // Left child index.
	right := 2*i + 2 // Right child index.

	// Check if the left child exists and is greater than the root.
	if left < len(arr) && arr[left] > arr[largest] {
		largest = left
	}
	// Check if the right child exists and is greater than the current largest.
	if right < len(arr) && arr[right] > arr[largest] {
		largest = right
	}
	// If largest is not root, swap it with the largest and continue heapifying.
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest)
	}
}

// buildMaxHeap converts an array into a max heap.
func buildMaxHeap(arr []int) {
	// Start from the last non-leaf node and move upwards.
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i)
	}
}

func main() {
	input := []int{4, 9, 16}
	k := 2
	fmt.Println(takeGiftsFromRichestMax(input, k))
}
