// See https://codility.com/media/train/3-PrefixSums.pdf

package main

import "fmt"

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	sLen := len(S)
	pLen := len(P)
	ps := make([][]int, 4)
	for i := 0; i < 4; i++ {
		ps[i] = make([]int, sLen+1)
	}

	for i := 0; i < sLen; i++ {
		c := S[i : i+1]
		if c == "A" {
			ps[0][i+1] = 1
		}
		if c == "C" {
			ps[1][i+1] = 1
		}
		if c == "G" {
			ps[2][i+1] = 1
		}
		if c == "T" {
			ps[3][i+1] = 1
		}
	}

	// Calculate prefix sums
	fmt.Println("-")
	for i := 0; i < 4; i++ {
		for j := 1; j < sLen+1; j++ {
			ps[i][j] += ps[i][j-1]
		}
		fmt.Println(ps[i])
	}

	result := make([]int, pLen)
	for i := 0; i < pLen; i++ {
		fromIndex := P[i]
		toIndex := Q[i] + 1

		if ps[0][toIndex]-ps[0][fromIndex] > 0 {
			result[i] = 1
		} else if ps[1][toIndex]-ps[1][fromIndex] > 0 {
			result[i] = 2
		} else if ps[2][toIndex]-ps[2][fromIndex] > 0 {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}
	return result
}
