package main

import (
	"reflect"
	"testing"
)

func TestGenomicRangeQuery(t *testing.T) {
	var tests = []struct {
		S   string
		P   []int
		Q   []int
		res []int
	}{
		{S: "CAGCCTA", P: []int{2, 5, 0}, Q: []int{4, 5, 6}, res: []int{2, 4, 1}},
		{S: "A", P: []int{0}, Q: []int{0}, res: []int{1}},
		{S: "AC", P: []int{0, 0, 1}, Q: []int{0, 1, 1}, res: []int{1, 1, 2}},
	}

	for _, test := range tests {
		result := GenomicRangeQuery(test.S, test.P, test.Q)
		if !reflect.DeepEqual(result, test.res) {
			t.Errorf("Expected %v, got %v", test.res, result)
		}
	}
}
