package main

func CyclicRotation(A []int, K int) []int {
	var B = make([]int, len(A))

	for i, val := range A {
		if i+K < len(A) {
			B[i+K] = val
		} else {
			B[(i+K)%len(A)] = val
		}
	}

	return B
}
