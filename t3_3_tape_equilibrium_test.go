package main

import "testing"

func TestTapeEquilibrium(t *testing.T) {
	var tests = []struct {
		desc string
		arr  []int
		res  int
	}{
		{"Single Element", []int{2}, 0},
		{"Two Elements", []int{2, 5}, 3},
		{"{3, 1, 2, 4, 3}", []int{3, 1, 2, 4, 3}, 1},
		{"{3, 1, 2, 4, 3}", []int{-3, -1, -2, -4, -3}, 1},
	}

	for _, test := range tests {
		result := TapeEquilibrium(test.arr)
		if result != test.res {
			t.Errorf("%v - Got: %d, want: %d.", test.desc, result, test.res)
		}
	}
}
