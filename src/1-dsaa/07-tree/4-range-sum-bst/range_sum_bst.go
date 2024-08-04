package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(H), where H is the hight of the bst,
// space is used by the call stack during the recursion.
func rangeSumBst(input *TreeNode, min int, max int) int {
	solution := Solution{
		done: false,
		sum:  0,
	}

	inorderTraversalWithSkipRanges(input, min, max, &solution)
	return solution.sum
}

type Solution struct {
	done bool
	sum  int
}

func inorderTraversalWithSkipRanges(input *TreeNode, min int, max int, solution *Solution) {
	if input == nil {
		return
	}

	if solution.done {
		return
	}

	inorderTraversalWithSkipRanges(input.Left, min, max, solution)

	if input.Value >= min && input.Value <= max {
		solution.sum += input.Value
	} else if input.Value > max {
		solution.done = true
	}

	if solution.done {
		return
	}

	inorderTraversalWithSkipRanges(input.Right, min, max, solution)
}

func main() {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Value: 3,
			},
			Right: &TreeNode{
				Value: 7,
			},
			Value: 5,
		},
		Right: &TreeNode{
			Right: &TreeNode{
				Value: 18,
			},
			Value: 15,
		},
		Value: 10,
	}
	min := 7
	max := 15
	fmt.Println(rangeSumBst(input, min, max))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
