package main

import "testing"

func TestMissingInteger(t *testing.T) {
	var tests = []struct {
		arr []int
		res int
	}{
		{arr: []int{1, 3, 4}, res: 2},
	}

	for _, test := range tests {
		result := MissingInteger(test.arr)
		if result != test.res {
			t.Errorf("Expected %d, got %d", test.res, result)
		}
	}
}
