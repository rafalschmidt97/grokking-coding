package main

import (
	"container/list"
	"strings"
)

func simplifyPathContainers(input string) string {
	splittedInput := strings.Split(input, "/")
	stack := list.New() // Use list as a stack

	// Filter components and build the stack
	for _, v := range splittedInput {
		if v == "" || v == "." {
			continue
		} else if v == ".." {
			// If "..", pop from the stack if it's not empty
			if stack.Len() > 0 {
				stack.Remove(stack.Back())
			}
		} else {
			// Push the component onto the stack
			stack.PushBack(v)
		}
	}

	// Build the simplified path
	var result strings.Builder
	for e := stack.Front(); e != nil; e = e.Next() {
		result.WriteString("/")
		result.WriteString(e.Value.(string))
	}

	// If the result is empty, return root "/"
	if result.Len() == 0 {
		return "/"
	}
	return result.String()
}
