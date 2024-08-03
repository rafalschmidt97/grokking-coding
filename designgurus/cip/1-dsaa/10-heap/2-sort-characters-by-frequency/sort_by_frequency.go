package main

import (
	"fmt"
	"strings"
)

// Arrange in desc order based on frequency of char
// Input: "apple"
// Expected Output: "ppale" or "ppela"
//
// Time complexity: O(n m logn m) = O(nmlogm)
// Space complexity: O(1 m) = O(m)
func sortByFrequency(input string) string {
	frequencies := make(map[rune]int, 0)
	for _, v := range input {
		frequencies[v]++
	}

	maxHeap := make([]rune, 0)
	for _, v := range input {
		maxHeap = append(maxHeap, v)
	}
	buildMaxHeap(maxHeap, frequencies)

	var builder strings.Builder
	// for _, v := range maxHeap { // cannot walk through
	for len(maxHeap) > 0 {
		v := maxHeap[0] // peak
		builder.WriteRune(v)
		maxHeap = maxHeap[1:] // pop
		buildMaxHeap(maxHeap, frequencies)
	}
	return builder.String()
}

func main() {
	input := "apple"
	fmt.Println(sortByFrequency(input))
}

// Boilerplate

// ensures the max heap property for the subtree rooted at index i.
func maxHeapify(arr []rune, i int, frequencies map[rune]int) {
	largest := i     // Initialize largest as root.
	left := 2*i + 1  // Left child index.
	right := 2*i + 2 // Right child index.

	// Check if the left child exists and is greater than the root.
	if left < len(arr) && frequencies[arr[left]] > frequencies[arr[largest]] {
		largest = left
	}
	// Check if the right child exists and is greater than the current largest.
	if right < len(arr) && frequencies[arr[right]] > frequencies[arr[largest]] {
		largest = right
	}
	// If largest is not root, swap it with the largest and continue heapifying.
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest, frequencies)
	}
}

// converts an array into a max heap.
func buildMaxHeap(arr []rune, frequencies map[rune]int) {
	// Start from the last non-leaf node and move upwards.
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, frequencies)
	}
}
