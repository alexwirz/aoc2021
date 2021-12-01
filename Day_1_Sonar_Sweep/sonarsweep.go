package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
