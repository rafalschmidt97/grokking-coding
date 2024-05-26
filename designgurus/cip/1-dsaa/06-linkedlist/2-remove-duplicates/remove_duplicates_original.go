package main

// Input is sorted.
// Time complexity: O(n)
// Space complexity: O(1):
func removeDuplicatesOriginal(input *ListNode) *ListNode {
	lastNode := input

	for lastNode != nil && lastNode.Next != nil {
		if lastNode.Next.Value == lastNode.Value {
			lastNode.Next = lastNode.Next.Next // bypass
		} else {
			lastNode = lastNode.Next
		}
	}

	return input
}
