package main

import (
	"fmt"
	"slices"
)

// Time complexity: O(n)
// Space complexity: O(n)
func verifyBalancedParentheses(input string) bool {
	openingSymbols := []rune{'(', '[', '{'}
	closingSymbols := []rune{')', ']', '}'}

	stack := make([]rune, 0, len(input)) // could use 'byte' as well
	position := -1                       // not initialized

	for _, v := range input {
		if !slices.Contains(openingSymbols, v) && !slices.Contains(closingSymbols, v) {
			panic("Unexpected symbol")
		}

		if position == -1 {
			if slices.Contains(closingSymbols, v) {
				return false
			}
		}

		if slices.Contains(openingSymbols, v) {
			stack = append(stack, v)
			position = position + 1
			continue
		}

		if idx := slices.Index(closingSymbols, v); idx != -1 {
			if stack[position] == openingSymbols[idx] {
				stack = slices.Delete(stack, position, position+1) // pop, O(1)
				position = position - 1
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	input := "{[()]}"
	fmt.Println(verifyBalancedParentheses(input))
}
