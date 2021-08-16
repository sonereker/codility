package main

import (
	"math"
	"sort"
)

//TapeEquilibrium O(N)
func TapeEquilibrium(A []int) int {
	if len(A) == 1 {
		return 0
	}

	var m1 = make([]int, len(A)-1)

	var arraySum int
	for _, i2 := range A {
		arraySum += i2
	}

	var s1 int
	for i := 0; i < len(A)-1; i++ {
		s1 += A[i]
		s2 := arraySum - s1

		m1[i] = int(math.Abs(float64(s1 - s2)))
	}

	sort.Ints(m1)
	return m1[0]
}
