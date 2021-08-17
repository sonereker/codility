package main

func FrogRiver(X int, A []int) int {
	var requiredSteps = make(map[int]int, X)
	for i, val := range A {
		if val <= X {
			requiredSteps[val] = 1
		}

		if len(requiredSteps) == X {
			return i
		}
	}

	return -1
}
