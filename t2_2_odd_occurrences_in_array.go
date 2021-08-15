package main

func OddOccurrencesInArray(A []int) int {
	var odd int
	unique := make(map[int]int)

	for _, val := range A {
		unique[val]++
	}

	for key, value := range unique {
		if value%2 == 1 {
			odd = key
			break
		}
	}

	return odd
}
