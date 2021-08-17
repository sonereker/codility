package main

func FrogRiver(X int, A []int) int {
	var requiredSteps []int
	for i := 1; i < X+1; i++ {
		requiredSteps = append(requiredSteps, i)
	}

	for i, val := range A {
		if val <= X {
			for k, step := range requiredSteps {
				if step == val {
					requiredSteps = append(requiredSteps[:k], requiredSteps[k+1:]...)
				}
			}
		}

		if len(requiredSteps) == 0 {
			return i
		}
	}

	return -1
}
