package main

import "container/list"

func zigzagIteratorContainers(input [2][]int) []int {
	queue := list.New() // Use list as a queue

	v1, v2 := input[0], input[1]

	i, j := 0, 0
	for i < len(v1) || j < len(v2) {
		if i < len(v1) {
			queue.PushBack(v1[i])
			i++
		}
		if j < len(v2) {
			queue.PushBack(v2[j])
			j++
		}
	}

	// Collect the elements into a result slice
	result := make([]int, 0, queue.Len())
	for e := queue.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(int))
	}

	return result
}
