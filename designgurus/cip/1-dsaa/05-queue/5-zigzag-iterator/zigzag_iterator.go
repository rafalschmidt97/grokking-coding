package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(n)
func zigzagIterator(input [2][]int) []int {
	queue := []int{}

	vectorOne := input[0]
	vectorTwo := input[1]
	hasMoreVectorOne := len(vectorOne) > 0
	hasMoreVectorTwo := len(vectorTwo) > 0

	for hasMoreVectorOne || hasMoreVectorTwo {
		if hasMoreVectorOne {
			queue = append(queue, vectorOne[0])
			vectorOne = vectorOne[1:]
			hasMoreVectorOne = len(vectorOne) > 0
		}

		if hasMoreVectorTwo {
			queue = append(queue, vectorTwo[0])
			vectorTwo = vectorTwo[1:]
			hasMoreVectorTwo = len(vectorTwo) > 0
		}
	}

	return queue
}

func main() {
	input := [2][]int{
		{1, 2},
		{3, 4, 5, 6},
	}
	fmt.Println(zigzagIterator(input))
}
