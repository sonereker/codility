package main

import "testing"

func TestFrogRiver(t *testing.T) {
	var tests = []struct {
		arr []int
		pos int
		res int
	}{
		{arr: []int{1, 3, 1, 4, 2, 3, 5, 4}, pos: 5, res: 6},
		{arr: []int{1, 3, 1, 3, 2, 1, 3}, pos: 3, res: 4},
		{arr: []int{2, 2, 2, 2, 2}, pos: 2, res: -1},
	}

	for _, test := range tests {
		result := FrogRiver(test.pos, test.arr)
		if result != test.res {
			t.Errorf("Expected %d. got %d: ", test.res, result)
		}
	}
}
