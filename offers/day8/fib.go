package day8

// Fibonacci 暴力递归
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

var dic = make([]int, 0)

// Fibonacci2 递归+缓存数组
func Fibonacci2(n int) int {
	if n <= 1 {
		return n
	}
	if v := dic[n-2]; v != 0 {
		return dic[n-2]
	}

	res := Fibonacci2(n-2) + Fibonacci2(n-1)
	dic = append(dic, res)

	return res
}

// Fibonacci3 first+second
func Fibonacci3(n int) int {
	if n <= 1 {
		return n
	}

	first := 0
	second := 1
	for i := 0; i < n-1; i++ {
		sum := first + second
		first = second
		second = sum
	}

	return second
}

// Fibonacci4 动态规划
func Fibonacci4(n int) int {
	if n <= 1 {
		return n
	}

	first := 0
	second := 1
	sum := 1

	for i := 3; i <= n; i++ {
		first = second
		second = sum
		sum = first + second
	}

	return sum
}
