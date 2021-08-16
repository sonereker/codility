package main

import "testing"

func TestFrogJump(t *testing.T) {
	result := FrogJump(1, 5, 2)
	if result != 2 {
		t.Errorf("Got: %d, want: %d.", result, 2)
	}
}
