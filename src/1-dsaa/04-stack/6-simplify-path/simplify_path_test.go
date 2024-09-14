package main

import (
	"fmt"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "/a//b////c/d//././/..",
			expected: "/a/b/c",
		},
		{
			input:    "/../",
			expected: "/",
		},
		{
			input:    "/../..",
			expected: "/",
		},
		{
			input:    "/home//foo/",
			expected: "/home/foo",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := simplifyPathContainers(tt.input)

			if got != tt.expected {
				t.Errorf("got '%v'; expected '%v'", got, tt.expected)
			}
		})
	}
}
