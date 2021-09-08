package main

import "testing"

func TestPassingCars(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{name: "[0, 1, 0, 1, 1]", arr: []int{0, 1, 0, 1, 1}, want: 5},
		{name: "[1]", arr: []int{1}, want: 0},
		{name: "[1, 0]", arr: []int{1, 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PassingCars(tt.arr); got != tt.want {
				t.Errorf("PassingCars() = %v, want %v", got, tt.want)
			}
		})
	}
}
