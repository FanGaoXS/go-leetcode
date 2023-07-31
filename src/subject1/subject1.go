package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	index := twoSum(nums, target)
	fmt.Println(index[0])
	fmt.Println(index[1])
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
