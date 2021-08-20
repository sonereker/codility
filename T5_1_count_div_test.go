package main

import "testing"

func TestCountDiv(t *testing.T) {
	var tests = []struct {
		A   int
		B   int
		K   int
		res int
	}{
		{A: 11, B: 345, K: 17, res: 20},
		{A: 10, B: 10, K: 5, res: 1},
		{A: 0, B: 22, K: 3, res: 7},
		{A: 6, B: 11, K: 2, res: 3},
		{A: 1, B: 2, K: 2, res: 1},
		{A: 1, B: 2000_000_000, K: 2, res: 1000000000},
	}

	for _, test := range tests {
		result := CountDiv(test.A, test.B, test.K)
		if result != test.res {
			t.Errorf("Expected %d, got %d", test.res, result)
		}
	}
}
