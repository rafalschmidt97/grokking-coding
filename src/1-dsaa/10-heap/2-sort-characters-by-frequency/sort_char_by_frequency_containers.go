package main

import (
	"container/heap"
	"strings"
)

// Time complexity: O(n log m)
// Space complexity: O(m)
func sortByFrequencyContainers(input string) string {
	frequencies := make(map[rune]int)
	for _, v := range input {
		frequencies[v]++
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for char, freq := range frequencies {
		heap.Push(maxHeap, CharFrequency{char: char, frequency: freq})
	}

	var builder strings.Builder
	for maxHeap.Len() > 0 {
		charFreq := heap.Pop(maxHeap).(CharFrequency)
		for i := 0; i < charFreq.frequency; i++ {
			builder.WriteRune(charFreq.char)
		}
	}

	return builder.String()
}

// Boilerplate

// A CharFrequency struct that holds a character and its frequency.
type CharFrequency struct {
	char      rune
	frequency int
}

// MaxHeap implements heap.Interface and holds CharFrequency elements.
type MaxHeap []CharFrequency

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].frequency > h[j].frequency
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(CharFrequency))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
