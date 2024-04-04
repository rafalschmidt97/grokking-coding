package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{2, 3, 4, 5, 10, 40}
	x := 3
	want := 1

	got := binarySearch(array, x)

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
