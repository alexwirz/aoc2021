package main

import "fmt"

func CheckIncreased(measurements []int) int {
	var count int = 0
	var last int = int(^uint(0) >> 1)
	fmt.Printf("measurements: %v\n", measurements)
	for i, v := range measurements {
		if i == 0 {
			last = v
			continue
		}

		if v > last {
			fmt.Printf("%d increase (%d)\n", v, last)
			count++
		}

		last = v
	}

	return count
}

func CheckIncreased2(measurements []int) (count int) {
	if len(measurements) < 2 {
		return 0
	}

	var last int = measurements[0]
	fmt.Printf("measurements: %v\n", measurements)
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > last {
			fmt.Printf("%d increase (%d)\n", measurements[i], last)
			count++
		}

		last = measurements[i]
	}

	return count
}
