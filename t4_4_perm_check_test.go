package main

import "testing"

func TestPermCheck(t *testing.T) {
	var tests = []struct {
		arr []int
		res int
	}{
		{arr: []int{2}, res: 0},
		{arr: []int{1}, res: 1},
		{arr: []int{4, 1, 3, 2}, res: 1},
		{arr: []int{9, 5, 7, 3, 2, 7, 3, 1, 10, 8}, res: 0},
		{arr: []int{2, 2, 3, 2}, res: 0},
		{arr: []int{4, 1, 3}, res: 0},
	}

	for _, test := range tests {
		result := PermCheck(test.arr)
		if result != test.res {
			t.Errorf("%v - Expected %d, got %d", test.arr, test.res, result)
		}
	}
}
