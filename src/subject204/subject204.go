package main

import "math"

func main() {

}

func countPrimes(n int) int {
	sum := 0
	for i := 2; i < n; i++ {
		if isPrimeNumber(i) {
			sum++
		}
	}
	return sum
}

//是否是质数
func isPrimeNumber(n int) (ok bool)  {
	if n<=0 {	//0和负数都不是质数
		return false
	}
	if n <= 2 {
		return true
	}
	sqrt := math.Sqrt(float64(n)) //平方根
	for i := 2; i <= int(sqrt) ; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
