package main

import (
	"reflect"
	"testing"
)

func TestMaxCounters(t *testing.T) {
	var tests = []struct {
		arr   []int
		count int
		res   []int
	}{
		{arr: []int{3, 4, 4, 6, 1, 4, 4}, count: 5, res: []int{3, 2, 2, 4, 2}},
	}

	for _, test := range tests {
		result := MaxCounters(test.count, test.arr)
		if !reflect.DeepEqual(result, test.res) {
			t.Errorf("Expected %d, got %d", test.res, result)
		}
	}

}
