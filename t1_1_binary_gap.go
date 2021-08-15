package main

import (
	"fmt"
	"strconv"
)

func LongestBinaryGap(N int) int {
	n := int64(N)
	b2 := strconv.FormatInt(n, 2)

	var counter int
	var counts []int

	for i, v := range b2 {
		char := fmt.Sprintf("%c", v)
		if char == "1" {
			if counter != 0 {
				counts = append(counts, counter)
				counter = 0
			}
		} else {
			counter += 1
		}

		if len(b2) == i+1 && char == "1" && counter != 0 {
			counts = append(counts, counter)
		}
	}

	for _, count := range counts {
		if counts[0] < count {
			counts[0] = count
		}
	}

	if len(counts) > 0 {
		return counts[0]
	}

	return 0
}
