package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(n)
func removeStars(input string) string {
	stack := []rune{}

	for _, v := range input {
		if v == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

			// skip otherwise
		} else {
			stack = append(stack, v)
		}
	}

	return string(stack)
}

func main() {
	input := "abc*de*f"
	fmt.Println(removeStars(input))
}
