package main

import (
	"fmt"
)

// Time complexity: TODO:
// Space complexity: TODO:
func swapNodesPairs(input *ListNode) *ListNode {
	return input
}

func main() {
	input := []int{3, 5, 2}
	fmt.Println(swapNodesPairs(arrayToListNode(input)))
}

// Boilerplate for singly linked list

type ListNode struct {
	Next  *ListNode
	Value int
}

func arrayToListNode(input []int) *ListNode {
	var headNode *ListNode
	var previousNode *ListNode

	for _, v := range input {
		nextNode := ListNode{
			Value: v,
			Next:  nil,
		}

		if headNode == nil {
			headNode = &nextNode
		} else {
			previousNode.Next = &nextNode
		}

		previousNode = &nextNode
	}
	return headNode
}

func listNodeToArray(input *ListNode) []int {
	array := []int{}
	for e := input; e != nil; e = e.Next {
		array = append(array, e.Value)
	}
	return array
}
