package main

func rangeSumBstOriginal(input *TreeNode, min int, max int) int {
	if input == nil { // base
		return 0
	}

	if input.Value < min {
		return rangeSumBstOriginal(input.Right, min, max)
	}

	if input.Value > max {
		return rangeSumBstOriginal(input.Left, min, max)
	}

	return input.Value +
		rangeSumBstOriginal(input.Right, min, max) +
		rangeSumBstOriginal(input.Left, min, max)
}
