package main

import "sort"

//PermMissingElement O(N) or O(N * log(N))
func PermMissingElement(A []int) int {
	sort.Ints(A)
	for i, val := range A {
		//First element
		if i == 0 && val > 1 {
			return 1
		}

		//Last element
		if i+1 == len(A) {
			return val + 1
		}

		if A[i+1] != val+1 {
			return val + 1
		}
	}
	return 1 //Empty array
}
