package main

import "fmt"

func main() {
	n := 4
	res := climbStairs(n)
	fmt.Println("res =",res)
}

//斐波那契数列
func climbStairs(n int) int {
	if n <= 0 {
		return -1
	}
	if n <= 3 {
		return n
	}
	first := 1
	second := 2
	for i := 0; i < n-2; i++ {
		sum := first + second
		first = second
		second = sum
	}
	return second
}
