package main

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	array := []int{3, 5, 7, 9, 11}
	x := 7
	want := 2

	got, err := linearSearch(array, x)

	if got != want || err != nil {
		t.Errorf("got %v; want %v (err: %v)", got, want, err)
	}
}
