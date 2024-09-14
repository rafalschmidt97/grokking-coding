package main

import (
	"container/list"
	"strings"
)

func makeGreatContainers(input string) string {
	stack := list.New() // Use a linked list as the stack

	for _, v := range input {
		if stack.Len() > 0 {
			top := stack.Back().Value.(rune)
			// Check if the current character and the top character are the same (case-insensitive) but different in case
			if strings.EqualFold(string(top), string(v)) && top != v {
				stack.Remove(stack.Back()) // Pop the top character if the condition is met
				continue
			}
		}
		// Push the current character onto the stack
		stack.PushBack(v)
	}

	// Build the result string from the stack
	var result []rune
	for e := stack.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(rune))
	}

	return string(result)
}
