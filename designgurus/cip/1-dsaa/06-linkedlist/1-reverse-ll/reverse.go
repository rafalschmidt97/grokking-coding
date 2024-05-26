package main

import (
	"fmt"
)

// Time complexity: O(n):
// Space complexity: O(1):
func reverse(input *ListNode) *ListNode {
	var previousNode *ListNode
	currentNode := input

	for currentNode != nil {
		nextNode := currentNode.Next // store a pointer

		currentNode.Next = previousNode // reverse direction

		previousNode = currentNode // save pointer for the next loop iternation
		currentNode = nextNode     // save pointer for the next loop iternation
	}

	return previousNode
}

func main() {
	input := []int{3, 5, 2}
	fmt.Println(reverse(arrayToListNode(input)))
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
