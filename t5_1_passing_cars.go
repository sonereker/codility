package main

func PassingCars(A []int) int {
	arr := make([][]int, 2)
	for i := 0; i < 2; i++ {
		arr[i] = make([]int, len(A)+1)
	}

	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			arr[0][i+1] = 1
		}
		if A[i] == 1 {
			arr[1][i+1] = 1
		}
	}

	for i := 1; i < len(A); i++ {
		if A[i] == 0 {
			arr[0][i+1] += arr[0][i]
			arr[1][i+1] = arr[1][i]
		}

		if A[i] == 1 {
			arr[0][i+1] = arr[0][i]
			arr[1][i+1] += arr[1][i]
		}
	}

	sum := arr[0][len(A)] + arr[1][len(A)]

	if sum < 2 {
		return 0
	}

	if sum > 1_000_000_000 {
		return -1
	}

	return sum
}
