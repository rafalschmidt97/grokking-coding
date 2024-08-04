package main

import (
	"fmt"
)

// Input is sorted.
// Time complexity: O(n)
// Space complexity: O(1):
func removeDuplicates(input *ListNode) *ListNode {
	headNode := &ListNode{
		Value: input.Value,
		Next:  nil,
	}
	lastNode := headNode

	for e := input; e != nil; e = e.Next {
		if lastNode.Value != e.Value {
			newNode := &ListNode{
				Value: e.Value,
				Next:  nil,
			}

			lastNode.Next = newNode
			lastNode = newNode
		}
	}

	return headNode
}

func main() {
	input := []int{3, 5, 2}
	fmt.Println(removeDuplicates(arrayToListNode(input)))
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
