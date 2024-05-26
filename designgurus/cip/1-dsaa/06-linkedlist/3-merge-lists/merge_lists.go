package main

import (
	"fmt"
)

// Inputs are sorted.
// Time complexity: O(m+n)
// Space complexity: O(1):
func mergeLists(inputOne *ListNode, inputTwo *ListNode) *ListNode {
	head := &ListNode{
		Next:  nil,
		Value: 0, // assumes that the values are positive
	}
	current := head

	for inputOne != nil && inputTwo != nil {
		if inputOne.Value > inputTwo.Value {
			current.Next = inputTwo
			inputTwo = inputTwo.Next
		} else {
			current.Next = inputOne
			inputOne = inputOne.Next
		}

		current = current.Next
	}

	// Append remaining
	if inputOne != nil {
		current.Next = inputOne
	}

	if inputTwo != nil {
		current.Next = inputTwo
	}

	return head.Next // skip pointer head
}

func main() {
	inputOne := []int{
		1, 3, 5,
	}
	inputTwo := []int{
		2, 4, 6,
	}

	fmt.Println(mergeLists(arrayToListNode(inputOne), arrayToListNode(inputTwo)))
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
