package main

import (
	"fmt"
)

// Given a singly linked list, swap every two adjacent nodes and return the head of
// the modified list.
//
// If the total number of nodes in the list is odd, the last node remains in place.
// Every node in the linked list contains a single integer value.
//
// Time complexity: TODO:
// Space complexity: TODO:
func swapNodesPairs(input *ListNode) *ListNode {
	head := &ListNode{
		Next:  input,
		Value: 0,
	}
	previous := head // first one is dummy
	current := input

	for current != nil && current.Next != nil {
		first := current
		second := current.Next

		current.Next = second.Next // first node as second
		second.Next = first        // second node as first
		previous.Next = second

		current = first.Next // move on to the next pair
		previous = first
	}

	return head.Next // skip pointer head
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
