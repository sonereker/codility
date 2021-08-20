package main

import "math"

func CountDiv(A int, B int, K int) int {
	if A == B {
		if B%K == 0 {
			return 1
		} else {
			return 0
		}
	}

	res := int(Round(float64(B-A) / float64(K)))

	if A == 0 {
		res++
	}

	return res
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
