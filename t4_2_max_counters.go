package main

func MaxCounters(N int, A []int) []int {
	counter := make([]int, N)
	var max int

	for _, val := range A {
		if 1 <= val && val <= N {
			counter[val-1]++
			if counter[val-1] > max {
				max = counter[val-1]
			}
		}

		if val == N+1 {
			for k, _ := range counter {
				if counter[k] < max {
					counter[k] = max
				}
			}
		}
	}
	return counter
}
