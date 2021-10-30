package main

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	length := m+n
	nums := make([]int,length)
	for i := 0; i < length; i++ {
		if i>=m {
			nums[i] = nums2[i-m]
		} else {
			nums[i] = nums1[i]
		}
	}	//合并数组
	sort.Ints(nums)	//数组排序
	if length%2 == 0 {	//如果数组长度是偶数
		num1 := nums[length/2-1]
		num2 := nums[length/2]
		return (float64(num1)+float64(num2))/2
	}
	return float64(nums[length/2])
}
