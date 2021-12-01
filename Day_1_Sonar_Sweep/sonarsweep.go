package main

import "fmt"

func CheckIncreased(measurements []int) int {
	if len(measurements) < 2 {
		return 0
	}

	var count int = 0
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

func CheckIncreased2(measurements []int) int {
	if len(measurements) < 2 {
		return 0
	}

	var count int = 0
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
