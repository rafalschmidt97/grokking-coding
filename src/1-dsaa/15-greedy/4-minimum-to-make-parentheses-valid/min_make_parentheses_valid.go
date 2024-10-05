package main

import (
	"container/list"
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(n)
func minMakeParenthesesValidStack(input string) int {
	stack := list.New()
	missing := 0

	for _, v := range input {
		if v == ')' {
			if stack.Len() == 0 {
				missing++
				continue
			}

			if stack.Back().Value.(rune) == '(' {
				stack.Remove(stack.Back())
			} else {
				missing++
			}
		}

		if v == '(' {
			stack.PushBack(v)
			continue
		}
	}

	return missing + stack.Len()
}

// "))((", expected: 4
// "(()())(", expected: 1,
func main() {
	input := "(()" // expected 1
	fmt.Println(minMakeParenthesesValidStack(input))
}
