package day8

var mem = make([]int, 100)

func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if v := mem[n-2]; v != 0 {
		return v
	}

	res := (numWays(n-2) + numWays(n-1)) % 1000000007
	mem[n-2] = res

	return res
}

func numWays2(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	first := 1
	second := 1
	for i := 0; i < n-1; i++ {
		sum := (first + second) % 1000000007
		first = second
		second = sum
	}

	return second
}
