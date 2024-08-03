package main

import "fmt"

// Time complexity: TODO: change
// Space complexity: TODO: change
func findMedian(inputStream []int) []float64 {
	output := make([]float64, 0)
	maxHeap := make([]int, 0)
	buildMaxHeap(maxHeap)

	for _, v := range inputStream {
		if v == -1 {
			median := findMedianFromMaxHeap(maxHeap)
			output = append(output, median)
		} else {
			maxHeap = append(maxHeap, v)
			maxHeapify(maxHeap, 0)
		}
	}

	return output
}

func findMedianFromMaxHeap(maxHeap []int) float64 {
	i := len(maxHeap) / 2
	if len(maxHeap)%2 == 0 && len(maxHeap) > 0 {
		return float64(maxHeap[i]) + float64(maxHeap[i+1])/2
	} else {
		return float64(maxHeap[i])
	}
}

func main() {
	input := []int{3, 1, -1, 5, -1, 4, -1}
	fmt.Println(findMedian(input))
}

// Boilerplate

// ensures the max heap property for the subtree rooted at index i.
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

// converts an array into a max heap.
func buildMaxHeap(arr []int) {
	// Start from the last non-leaf node and move upwards.
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i)
	}
}
