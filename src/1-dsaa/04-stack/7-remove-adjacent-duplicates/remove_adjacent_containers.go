package main

import "container/list"

func removeAdjacentContainers(input string) string {
	stack := list.New() // Use list as the stack

	for _, v := range input {
		if stack.Len() == 0 || stack.Back().Value.(rune) != v {
			stack.PushBack(v) // Push the character onto the stack
		} else {
			stack.Remove(stack.Back()) // Pop the last element if it matches the current one
		}
	}

	// Build the result string from the stack
	var result []rune
	for e := stack.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(rune))
	}

	return string(result)
}
