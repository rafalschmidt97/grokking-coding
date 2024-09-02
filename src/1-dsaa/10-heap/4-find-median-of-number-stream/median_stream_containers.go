package main

import "container/heap"

// Time complexity: Insert: O(log n), Find: O(1)
// Space complexity: O(n)
func findMedianContainers(inputStream []int) []float64 {
	output := make([]float64, 0)
	lowerHalf := &MinHeap{} // Simulates a max-heap by storing negative values
	upperHalf := &MinHeap{}
	heap.Init(lowerHalf)
	heap.Init(upperHalf)

	for _, v := range inputStream {
		if v == -1 { // Code for find is minus 1
			// Find median
			if lowerHalf.Len() == upperHalf.Len() {
				median := float64(-(*lowerHalf)[0]+(*upperHalf)[0]) / 2
				output = append(output, median)
			} else {
				median := float64(-(*lowerHalf)[0])
				output = append(output, median)
			}
		} else {
			// Insert
			if lowerHalf.Len() == 0 || v <= -(*lowerHalf)[0] {
				heap.Push(lowerHalf, -v) // Push negative to simulate max-heap
			} else {
				heap.Push(upperHalf, v)
			}

			// Balance the heaps
			if lowerHalf.Len() > upperHalf.Len()+1 {
				heap.Push(upperHalf, -heap.Pop(lowerHalf).(int))
			} else if lowerHalf.Len() < upperHalf.Len() {
				heap.Push(lowerHalf, -heap.Pop(upperHalf).(int))
			}
		}
	}

	return output
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
