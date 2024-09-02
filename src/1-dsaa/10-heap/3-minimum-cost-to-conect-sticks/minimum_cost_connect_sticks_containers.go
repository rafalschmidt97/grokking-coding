package main

import "container/heap"

// Time complexity: O(n log n)
// Space complexity: O(n)
func minimumCostToConnectSticksContainers(input []int) int {
	minHeap := &MinHeap{}
	for _, v := range input {
		heap.Push(minHeap, v)
	}
	heap.Init(minHeap)

	costToConnect := 0

	// Continue combining the smallest two sticks until only one stick remains
	for minHeap.Len() > 1 {
		// Pop the two smallest sticks
		v1 := heap.Pop(minHeap).(int)
		v2 := heap.Pop(minHeap).(int)
		combined := v1 + v2

		// Add the cost of this combination
		costToConnect += combined

		// Push the combined stick back into the heap
		heap.Push(minHeap, combined)
	}

	return costToConnect
}

// Boilerplate

// MinHeap implements heap.Interface and holds ints.
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
