package main

import "math"

func CountDiv(A int, B int, K int) int {
	return int(Round(float64((B + 1 - A) / K)))
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
