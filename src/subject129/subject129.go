package main

import (
	. "go-leetcode/src/datastruct"
	"math"
)

var (
	sum int
	nums []int
)

func sumNumbers(root *TreeNode) int {
	sum = 0
	nums = []int{}
	dfs(root)
	return sum
}


func dfs(root *TreeNode)  {
	nums = append(nums,root.Val)
	//到达叶子节点
	if root.Left == nil && root.Right == nil {
		for i := 0; i < len(nums); i++ {
			num := nums[i]
			sum += num*int(math.Pow10(len(nums)-1-i))
		}
	}
	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
	nums = nums[:len(nums)-1]
}
