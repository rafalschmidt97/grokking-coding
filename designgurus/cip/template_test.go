package main

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	x := 1
	got := template(x)

	if got != 1 {
		t.Errorf("got %v; want %v", got, x)
	}
}

func TestTemplateCases(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{x: 1, want: 1},
		{x: -1, want: -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.x, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := template(tt.x)

			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
