package day8

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	first := 1
	second := 2
	for i := 1; i < n-1; i++ {
		sum := first + second
		first = second
		second = sum
	}

	return second
}
