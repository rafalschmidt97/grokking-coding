package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(n)
func removeAdjacent(input string) string {
	stack := []rune{}

	for _, v := range input {
		if len(stack) == 0 || stack[len(stack)-1] != v {
			stack = append(stack, v)
		} else {
			stack = stack[:len(stack)-1] // pop if matches
		}
	}

	return string(stack)
}

func main() {
	input := "abbaca"
	fmt.Println(removeAdjacent(input))
}
