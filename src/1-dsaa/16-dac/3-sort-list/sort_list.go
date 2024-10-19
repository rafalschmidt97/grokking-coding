package main

import "fmt"

// Time complexity: O(nlogn), splits and merges
// Space complexity: O(logn), recursive call stack
func sortList(head *ListNode) *ListNode {
	// Base case: a single node or empty list is already sorted
	if head == nil || head.Next == nil {
		return head
	}

	// Find the middle of the list to divide it into two halves
	// The slow pointer advances one node at a time, while the fast pointer advances two nodes at a time.
	// When the fast pointer reaches the end, the slow pointer will be at the midpoint.
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Disconnect the first half from the second half
	prev.Next = nil

	// Recursively sort each half
	l1 := sortList(head)
	l2 := sortList(slow)

	// Merge the two sorted halves and return
	return Merge(l1, l2)
}

// Merge - Merge two sorted linked lists
func Merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Value: 0}
	tail := dummy
	for l1 != nil && l2 != nil {
		// Compare values from each list and append the smaller one
		if l1.Value < l2.Value {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	// Append any remaining nodes from either list
	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}
	return dummy.Next // Return the head of the merged list
}

func main() {
	input := []int{1, 2, 3}
	fmt.Println(sortList(arrayToListNode(input)))
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
