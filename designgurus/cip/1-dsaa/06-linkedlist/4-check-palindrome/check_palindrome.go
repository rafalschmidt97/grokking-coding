package main

import (
	"fmt"
)

// A doubly linked list is a palindrome if it reads the same backward as forward,
// utilizing the previous and next pointers of the nodes.
//
// Time complexity: TODO:
// Space complexity: TODO:
func checkPalindrome(head *DoublyListNode) bool {
	tail := head

	for tail.Next != nil {
		tail = tail.Next
	}

	for tail != nil && head != nil {
		if tail.Value != head.Value {
			return false
		}

		head = head.Next
		tail = tail.Prev
	}

	return true
}

func main() {
	input := []int{1, 2, 3, 2, 1}
	fmt.Println(checkPalindrome(arrayToListNode(input)))
}

// Boilerplate for singly linked list

type DoublyListNode struct {
	Prev  *DoublyListNode
	Next  *DoublyListNode
	Value int
}

func arrayToListNode(input []int) *DoublyListNode {
	var headNode *DoublyListNode
	var previousNode *DoublyListNode

	for _, v := range input {
		nextNode := DoublyListNode{
			Value: v,
			Next:  nil,
			Prev:  nil,
		}

		if headNode == nil {
			headNode = &nextNode
		} else {
			previousNode.Next = &nextNode
			nextNode.Prev = previousNode
		}

		previousNode = &nextNode
	}
	return headNode
}
