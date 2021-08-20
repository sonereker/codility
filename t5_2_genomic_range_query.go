package main

import "sort"

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	var result []int

	dna := make(map[string]int)
	dna["A"] = 1
	dna["C"] = 2
	dna["G"] = 3
	dna["T"] = 4

	for i := 0; i <= len(P)-1; i++ {
		start := P[i]
		end := Q[i]

		var r []int
		for k := start; k <= end; k++ {
			r = append(r, dna[string(S[k])])
		}
		sort.Ints(r)
		result = append(result, r[0])
	}

	return result
}
