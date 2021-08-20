package main

import "math"

//FrogJump O(N)
func FrogJump(X int, Y int, D int) int {
	return int(HardRound(float64(Y-X) / float64(D)))
}

func HardRound(x float64) float64 {
	t := math.Trunc(x)
	if x != t {
		return t + math.Copysign(1, x)
	}
	return t
}
