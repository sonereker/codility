package main

import (
	"math"
	"sort"
)

//TapeEquilibrium O(N * N)
func TapeEquilibrium(A []int) int {
	if len(A) == 1 {
		return 0
	}

	var m1 = make([]int, len(A)-1)

	var arraySum int
	for _, i2 := range A {
		arraySum += i2
	}

	for i := 1; i < len(A); i++ {
		var a1 = A[0:i]

		var s1 int
		for _, v := range a1 {
			s1 += v
		}

		s2 := arraySum - s1

		m1[i-1] = int(math.Abs(float64(s1 - s2)))
	}

	sort.Ints(m1)
	return m1[0]
}
