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

func PartTwo(measurements []int) int {
	return 1
}

func Sum(numbers []int, start, end int) int {
	return 607
}

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func main() {
	nums, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(CheckIncreased(nums))
}
