package main

import "sort"

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	var result []int

	dna := make(map[string]int)
	dna["A"] = 1
	dna["C"] = 2
	dna["G"] = 3
	dna["T"] = 4

	var stringAsArray []int
	for i := 0; i < len(S); i++ {
		stringAsArray = append(stringAsArray, dna[string(S[i])])
	}

	for i := 0; i < len(P); i++ {
		start := P[i]
		end := Q[i]

		impactFactors := make([]int, len(stringAsArray[start:end+1]))
		copy(impactFactors, stringAsArray[start:end+1])
		sort.Ints(impactFactors)
		result = append(result, impactFactors[0])
	}

	return result
}
