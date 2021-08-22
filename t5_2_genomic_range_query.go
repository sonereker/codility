package main

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	var sequenceLength = len(S) + 1
	var genoms = make([][]int, 3)
	for i := range genoms {
		genoms[i] = make([]int, sequenceLength)
	}

	var a, c, g int
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "A" {
			a = 1
		}

		if string(S[i]) == "C" {
			c = 1
		}

		if string(S[i]) == "G" {
			g = 1
		}

		genoms[0][i+1] = genoms[0][i] + a
		genoms[1][i+1] = genoms[1][i] + c
		genoms[2][i+1] = genoms[2][i] + g
	}

	var result = make([]int, len(P))
	for i := 0; i < len(P); i++ {
		var fromIndex = P[i]
		var toIndex = Q[i] + 1

		if genoms[0][toIndex]-genoms[0][fromIndex] > 0 {
			result[i] = 1
		} else if genoms[1][toIndex]-genoms[1][fromIndex] > 0 {
			result[i] = 2
		} else if genoms[2][toIndex]-genoms[2][fromIndex] > 0 {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}

	return result
}
