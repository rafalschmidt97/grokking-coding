package main

import (
	"container/list"
	"fmt"
)

// Time complexity: TODO:
// Space complexity: TODO:
func reverse(input *list.List) *list.List {
	return input
}

func main() {
	input := []int{3, 5, 2}
	fmt.Println(reverse(arrayToLinkedList(input)))
}

func arrayToLinkedList(input []int) *list.List {
	linkedList := list.New()
	for _, v := range input {
		linkedList.PushBack(v)
	}
	return linkedList
}
