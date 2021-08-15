package main

import "testing"

func TestFrogJump(t *testing.T) {
	steps := FrogJump(1, 5, 2)
	if steps != 2 {
		t.Errorf("Got: %d, want: %d.", steps, 2)
	}
}
