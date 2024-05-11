package main

import (
	"fmt"
	"strings"
)

// Input: "/a//b////c/d//././/.."
// Input: a ' ' b ' ' ' ' ' ' c d . . .. --- split
// Input: a b c d . . .. --- remove empty
// Input: a b c d .. --- remove single dots
// Input: a b c --- go back
// Input: /a/b/c --- recreate from the stack
//
// Time complexity: O(n)
// Space complexity: O(n)
func simplifyPath(input string) string {
	splittedInput := strings.Split(input, "/")
	stack := make([]string, 0, len(splittedInput))

	// filter
	for _, v := range splittedInput {
		if v != "" && v != " " && v != "." {
			stack = append(stack, v)
		}
	}

	// go back if necessary
	i := len(stack) - 1
	for i >= 0 {
		if stack[i] == ".." {
			toPop := 0

			if len(stack) == 1 {
				toPop = 1
			} else {
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
