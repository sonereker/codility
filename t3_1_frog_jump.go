package main

import "math"

func FrogJump(X int, Y int, D int) int {
	return int(Round(float64(Y-X) / float64(D)))
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if x != t {
		return t + math.Copysign(1, x)
	}
	return t
}
