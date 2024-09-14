package main

import (
	"container/list"
)

// Time complexity: O(n)
// Space complexity: O(n)
func verifyBalancedParenthesesContainers(input string) bool {
	openingSymbols := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	closingSymbols := map[rune]rune{')': '(', ']': '[', '}': '{'}

	stack := list.New()

	for _, v := range input {
		if _, isOpening := openingSymbols[v]; isOpening {
			stack.PushBack(v)
		} else if matchingOpening, isClosing := closingSymbols[v]; isClosing {
			if stack.Len() == 0 || stack.Back().Value.(rune) != matchingOpening {
				return false
			}

			// Containers/list is a doubly linked list, which means it's not as optimized
			// for operations like "pop from the end" as a slice would be. Unfortunately,
			// there's no direct way to access or remove the last element
			// without using Remove in container/list.
			stack.Remove(stack.Back()) // Mimic the stack pop, O(n)
		} else {
			panic("Unexpected symbol")
		}
	}

	return stack.Len() == 0
}
