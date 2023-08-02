package offer11

import (
	"math"
	"sort"
)

func minArray(numbers []int) int {
	min := math.MaxInt
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}

	return min
}

func minArray2(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
