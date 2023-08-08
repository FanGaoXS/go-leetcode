package day8

var cache = make([]int, 101-2)

func fib(n int) int {
	if n <= 1 {
		return n
	}

	if cache[n-2] != 0 {
		return cache[n-2]
	}
	res := (fib(n-1) + fib(n-2)) % 1000000007
	cache[n-2] = res

	return res
}

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	p := 0
	q := 1
	r := 1
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % 1000000007
	}

	return r
}

func fib3(n int) int {
	if n <= 1 {
		return n
	}

	first := 0
	second := 1
	for i := 0; i < n; i++ {
		sum := (first + second) % 1000000007
		first = second
		second = sum
	}

	return first
}
