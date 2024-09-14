package main

import (
	"container/list"
	"strconv"
)

// Time complexity: O(logn) due to division by 2 at each step
// Space complexity: O(logn)
func decToBinContainers(input int) string {
	stack := list.New()
	leftAcum := input

	for leftAcum != 0 {
		modulo := leftAcum % 2 // Push the remainder (modulo 2) to the stack for each division
		stack.PushBack(modulo)
		leftAcum /= 2
	}

	var result string
	for stack.Len() > 0 {
		pop := stack.Back() // Pop the last element
		result += strconv.Itoa(pop.Value.(int))
		stack.Remove(pop) // Remove the last element (pop it)
	}
	return result
}
