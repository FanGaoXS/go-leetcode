package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10
	result := isPrimeNumber(n)
	fmt.Println("result =",result)
	nums := primeFactorization(n)
	fmt.Printf("nums = %v",nums)
	a := 1
	b := 3
	fmt.Println("c =",float64(a)/float64(b))
	
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

//分解质因数
func primeFactorization(n int) (nums []int)  {
	if n<=0 {
		return
	}
	for !isPrimeNumber(n) {
		sqrt := math.Sqrt(float64(n))
		for i := 2; i <= int(sqrt) ; i++ {
			if n%i == 0 {
				nums = append(nums,i)
				n = n/i
				break
			}
		}
	}
	nums = append(nums,n)
	return nums
}

func sumIsInt(a,b int) (ok bool) {
	return !(math.MaxInt - a < b)
}
