package main

//MaxCounters O(N + M)
func MaxCounters(N int, A []int) []int {
	counter := make([]int, N)
	var currentMax int
	var resetMax int

	for _, val := range A {
		if 1 <= val && val <= N {
			if resetMax != 0 && counter[val-1] < resetMax {
				counter[val-1] = resetMax
			}

			counter[val-1]++
			if counter[val-1] > currentMax {
				currentMax = counter[val-1]
			}
		}

		if val == N+1 {
			resetMax = currentMax
		}
	}

	if resetMax != 0 {
		for i, cval := range counter {
			if cval < resetMax {
				counter[i] = resetMax
			}
		}
	}

	return counter
}
