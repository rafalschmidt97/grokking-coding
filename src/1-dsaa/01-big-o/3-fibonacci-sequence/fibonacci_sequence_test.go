package main

import (
	"testing"
)

func TestFibonaciiSequence(t *testing.T) {
	n := 10
	expected := 55

	got := fibonacciSequence(n)

	if got != expected {
		t.Errorf("got %v; want %v for %v", got, expected, n)
	}
}
