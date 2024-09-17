package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Time complexity: O(nlogn) - for each meeting plus push and pop
// Space complexity: O(n)
func meetingsRoom(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort the meetings by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Create a min-heap to keep track of end times of meetings
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Add the first meeting's end time to the heap
	heap.Push(minHeap, intervals[0][1])

	// Process all other meetings
	for i := 1; i < len(intervals); i++ {
		// If the room with the earliest end time is free (the earliest end time is <= current meeting start time)
		if (*minHeap)[0] <= intervals[i][0] {
			// Remove the room from the heap
			heap.Pop(minHeap)
		}

		// Add the current meeting's end time to the heap
		heap.Push(minHeap, intervals[i][1])
	}

	// The size of the heap is the number of rooms needed
	return minHeap.Len()
}

func main() {
	input := [][]int{
		{10, 15},
		{20, 25},
		{30, 35},
	}
	fmt.Println(meetingsRoom(input))
}

// A min-heap to store the end times of the meetings
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
