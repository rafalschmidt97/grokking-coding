package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(n)
func removeStars(input string) string {
	stack := []rune{}
	for i, v := range input {
		if i == 0 {
			stack = append(stack, v)
		} else {
			if v == '*' {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1] // pop
				}
			} else {
				stack = append(stack, v)
			}
		}
	}
	return string(stack)
}

func main() {
	input := "abc*de*f"
	fmt.Println(removeStars(input))
}
