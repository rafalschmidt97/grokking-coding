package main

import (
	"container/heap"
)

type Entry struct {
	Num       int
	Frequency int
}

// MinHeapFrequencyEntries implements heap.Interface and holds Entries.
type MinHeapFrequencyEntries []Entry

// Time complexity: O(n log k)
// Space complexity: O(n + k)
func findTopKFrequentNumbersContainers(input []int, k int) []int {
	frequencies := make(map[int]int)
	for _, v := range input {
		frequencies[v]++
	}

	minHeap := &MinHeapFrequencyEntries{}
	heap.Init(minHeap)

	// Go through all numbers in numFrequencyMap and push them into the min-heap
	// If the heap size is more than k, remove the smallest (top) entry
	for num, frequency := range frequencies {
		entry := Entry{num, frequency}
		heap.Push(minHeap, entry)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	// Create a list of top k frequent numbers
	output := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		entry := heap.Pop(minHeap).(Entry)
		output[i] = entry.Num
	}

	return output
}

// Boilerplate
func (h MinHeapFrequencyEntries) Len() int {
	return len(h)
}

func (h MinHeapFrequencyEntries) Less(i, j int) bool { // min heap
	return h[i].Frequency < h[j].Frequency
}

func (h MinHeapFrequencyEntries) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapFrequencyEntries) Push(x interface{}) {
	*h = append(*h, x.(Entry))
}

func (h *MinHeapFrequencyEntries) Pop() interface{} {
	old := *h
	n := len(old)
	entry := old[n-1]
	*h = old[0 : n-1]
	return entry
}
