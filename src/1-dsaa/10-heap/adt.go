package main

import (
	"fmt"
)

type Solution struct{}

func (s *Solution) PrintMinHeap(arr []int, size int) {
	fmt.Print("Min Heap: ")
	for index := 0; index < size; index++ {
		fmt.Print(arr[index], " ")
	}
	fmt.Println()
}

func (s *Solution) Heapify(arr []int, size, index int) {
	smallestValue := index
	leftChild := 2*index + 1
	rightChild := 2*index + 2

	if leftChild < size && arr[leftChild] < arr[smallestValue] {
		smallestValue = leftChild
	}

	if rightChild < size && arr[rightChild] < arr[smallestValue] {
		smallestValue = rightChild
	}

	if smallestValue != index {
		temp := arr[index]
		arr[index] = arr[smallestValue]
		arr[smallestValue] = temp
		s.Heapify(arr, size, smallestValue)
	}
}

func (s *Solution) InsertMinHeap(arr []int, value, size int) int {
	arr[size] = value
	size++

	for index := size/2 - 1; index >= 0; index-- {
		s.Heapify(arr, size, index)
	}

	return size
}

func main() {
	minHeap := make([]int, 100) // Assuming max heap size is 100
	size := 0                   // Current size of the heap

	// Test case 1
	size = new(Solution).InsertMinHeap(minHeap, 5, size)
	new(Solution).PrintMinHeap(minHeap, size)

	// Test case 2
	size = new(Solution).InsertMinHeap(minHeap, 2, size)
	new(Solution).PrintMinHeap(minHeap, size)

	// Test case 3
	size = new(Solution).InsertMinHeap(minHeap, 8, size)
	new(Solution).PrintMinHeap(minHeap, size)
}
