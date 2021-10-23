package main

import "fmt"

func main() {
	nums := []int{2,5,4,9,1}
	result := findKthLargest(nums,5)
	fmt.Println("result =",result)
}

func findKthLargest(nums []int, k int) int {
	//冒泡排序（由大到小排序）
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			//交换
			if nums[j] < nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	return nums[k-1]
}


