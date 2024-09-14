package main

import "container/list"

func removeStarsContainers(input string) string {
	stack := list.New() // Use a linked list as the stack

	for _, v := range input {
		if v == '*' {
			// Pop the last character from the stack if it exists
			if stack.Len() > 0 {
				stack.Remove(stack.Back())
			}
		} else {
			// Push the current character onto the stack
			stack.PushBack(v)
		}
	}

	// Build the result string from the stack
	var result []rune
	for e := stack.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(rune))
	}

	return string(result)
}
