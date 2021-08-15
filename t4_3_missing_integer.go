package main

import (
	"sort"
)

func MissingInteger(A []int) int {
	A = append(A, 0)
	sort.Ints(A)

	for i, val := range A {
		if val < 1 {
			A[i] = 0
			val = 0
		}

		if len(A) == i+1 {
			return val + 1
		}

		if A[i+1] < 0 {
			continue
		}

		if A[i+1] != val+1 && A[i+1] != val {
			return val + 1
		}
	}

	return 1
}
