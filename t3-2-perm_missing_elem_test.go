package main

import "testing"

func TestPermMissingElement(t *testing.T) {
	arr := []int{2, 3, 1, 5}
	result := PermMissingElement(arr)
	if result != 4 {
		t.Errorf("Got: %d, want: %d.", result, 4)
	}
}
