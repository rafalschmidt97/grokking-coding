package main

import (
	"container/heap"
	"math"
)

// Time complexity: O(k log n + n)
// Space complexity: O(n)
func takeGiftsFromRichestMaxContainers(input []int, k int) int {
	maxHeap := &MaxHeap{}
	for _, v := range input {
		heap.Push(maxHeap, v)
	}
	heap.Init(maxHeap)

	// Perform k iterations of taking gifts
	for i := 0; i < k; i++ {
		// Extract the maximum value from the heap
		max := heap.Pop(maxHeap).(int)

		// Calculate the square root and push it back into the heap
		maxSqrt := int(math.Floor(math.Sqrt(float64(max))))
		heap.Push(maxHeap, maxSqrt)
	}

	// Calculate the remaining sum of the heap
	remaining := 0
	for maxHeap.Len() > 0 {
		remaining += heap.Pop(maxHeap).(int)
	}
	return remaining
}

// Boilerplate

// MaxHeap implements heap.Interface and holds ints.
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
