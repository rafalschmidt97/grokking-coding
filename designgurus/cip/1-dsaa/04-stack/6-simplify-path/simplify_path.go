package main

import (
	"fmt"
	"strings"
)

// Input: "/a//b////c/d//././/.."
// Input: "/a/b/c/d/././.." --- clean double shashes
// Input: "/a/b/c/d/.." -- clean singular dots
// Input: "/a/b/c" -- go back
//
// Time complexity: TODO:
// Space complexity: TODO:
func simplifyPath(input string) string {
	// inputStack := make([]rune, 0, len(input))
	outputStack := make([]rune, 0, len(input))

	// for _, v := range inputStack {
	// 	for len(outputStack) > 0 {
	// 		stackPeek := outputStack[len(outputStack)-1]
	// 		outputStack = outputStack[:len(outputStack)-1] // pop
	// 	}
	//
	// 	stack = append(stack, v)
	// }

	var builder strings.Builder
	for _, v := range outputStack {
		builder.WriteRune(v)
	}
	return builder.String()
}

func main() {
	input := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(input))
}
