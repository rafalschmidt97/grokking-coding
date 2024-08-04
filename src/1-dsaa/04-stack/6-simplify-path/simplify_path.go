package main

import (
	"fmt"
	"strings"
)

// Input: "/a//b////c/d//././/.."
// Input: ["", "a", "", "b", "", "", "", "c", "d", "", ".", "", ".", "", "", ".."]
// Input: none (ignored)
// Input: a
// Input: a - none (ignored)
// Input: a b c d
// Input: a b c
// Input: /a/b/c
//
// Time complexity: O(n)
// Space complexity: O(n)
func simplifyPath(input string) string {
	splittedInput := strings.Split(input, "/")
	stack := make([]string, 0, len(splittedInput))

	// filter first
	for _, v := range splittedInput {
		if v != "" && v != "." {
			// If the component is not empty and not '.', push it onto the stack
			stack = append(stack, v)
		}
	}

	// go back if necessary
	i := len(stack) - 1
	for i >= 0 {
		if stack[i] == ".." {
			toPop := 0

			if len(stack) == 1 {
				// If the component is '..', pop the the only one from the stack
				toPop = 1
			} else {
				// If the component is '..', pop the last two from the stack
				toPop = 2
			}

			stack = stack[:len(stack)-toPop] // pop
			i = i - toPop
		}

		i--
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	input := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(input))
}
