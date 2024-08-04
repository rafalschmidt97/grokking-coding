package main

import (
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
func simplifyPathOriginal(input string) string {
	var stack []string

	for _, p := range strings.Split(input, "/") {
		switch {
		case p == "..":
			// If the component is '..', pop the last component from the stack
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case p != "" && p != ".":
			// If the component is not empty and not '.', push it onto the stack
			stack = append(stack, p)
		}
	}

	// Reconstruct the simplified path by joining components from the stack
	return "/" + strings.Join(stack, "/")
}
