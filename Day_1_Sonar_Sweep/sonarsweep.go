package main

import "fmt"

func CheckIncreased(measurements []int) int {
	var count int = 0
	var last int = int(^uint(0) >> 1)
	for i, v := range measurements {
		if i == 0 {
			last = v
			continue
		}

		if v > last {
			fmt.Printf("%d increase (%d)\n", v, last)
			last = v
			count++
		}
	}

	return count
}
