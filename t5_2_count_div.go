package main

//CountDiv O(1)
func CountDiv(A int, B int, K int) int {
	if A == B {
		if B%K == 0 {
			return 1
		} else {
			return 0
		}
	}

	res := B/K - A/K

	if A%K == 0 {
		res++
	}

	return res
}
