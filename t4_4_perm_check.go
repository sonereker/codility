package main

import "sort"

//PermCheck O(N) or O(N * log(N))
func PermCheck(A []int) int {
	sort.Ints(A)

	if A[0] != 1 {
		return 0
	}

	for i, val := range A {
		if i != val-1 {
			return 0
		}
	}

	if A[len(A)-1] == len(A) {
		return 1
	}
	return 0
}
