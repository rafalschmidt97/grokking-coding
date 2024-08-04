package main

import "fmt"

type Solution struct {
	counter int
	result  int
}

// Time complexity: O(n) worst case, O(k) best case
// Space complexity: O(1) if stack not counts
func kthSmallestElementBst(input *TreeNode, k int) int {
	solution := Solution{
		counter: 0,
		result:  -1,
	}

	inorderTraversalWithStopOnFindK(input, k, &solution)
	return solution.result
}

func inorderTraversalWithStopOnFindK(input *TreeNode, k int, solution *Solution) {
	if input == nil {
		return
	}

	if solution.counter == k {
		return
	}

	inorderTraversalWithStopOnFindK(input.Left, k, solution)

	solution.counter++

	if solution.counter == k {
		solution.result = input.Value
	}

	inorderTraversalWithStopOnFindK(input.Right, k, solution)
}

func main() {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Value: 1,
			},
			Right: &TreeNode{
				Left: &TreeNode{
					Value: 4,
				},
				Right: &TreeNode{
					Value: 7,
				},
				Value: 6,
			},
			Value: 3,
		},
		Right: &TreeNode{
			Right: &TreeNode{
				Left: &TreeNode{
					Value: 13,
				},
				Value: 14,
			},
			Value: 10,
		},
		Value: 8,
	}
	k := 4
	fmt.Println(kthSmallestElementBst(input, k))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
