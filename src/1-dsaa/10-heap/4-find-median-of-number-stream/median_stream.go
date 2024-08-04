package main

import "fmt"

// Time complexity: Insert: O(logn), Find: O(1)
// Space complexity: O(n)
func findMedian(inputStream []int) []float64 {
	output := make([]float64, 0)
	maxHeap := make([]int, 0)
	minHeap := make([]int, 0)

	for _, v := range inputStream {
		if v == -1 { // code for find is minus 1
			// find

			if len(maxHeap) == len(minHeap) {
				median := float64(-maxHeap[0]+minHeap[0]) / 2
				output = append(output, median)
			} else {
				median := float64(-maxHeap[0])
				output = append(output, median)
			}
		} else {
			// insert

			if len(maxHeap) == 0 || -maxHeap[0] >= v {
				maxHeap = append(maxHeap, -v)
				minHeapify(maxHeap, 0)
			} else {
				minHeap = append(minHeap, v)
				minHeapify(minHeap, 0)
			}

			// either both the heaps will have equal number
			// of elements or max-heap will have one
			// more element than the min-heap
			if len(maxHeap) > len(minHeap)+1 {
				pullMaxHeap := -maxHeap[0]
				maxHeap = maxHeap[1:]
				minHeapify(maxHeap, 0)
				minHeap = append(minHeap, pullMaxHeap)
				minHeapify(minHeap, 0)
			} else if len(maxHeap) < len(minHeap) {
				pullMinHeap := minHeap[0]
				minHeap = minHeap[1:]
				minHeapify(minHeap, 0)
				maxHeap = append(maxHeap, -pullMinHeap)
				minHeapify(maxHeap, 0)
			}
		}
	}

	return output
}

func main() {
	input := []int{3, 1, -1, 5, -1, 4, -1}
	fmt.Println(findMedian(input))
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
