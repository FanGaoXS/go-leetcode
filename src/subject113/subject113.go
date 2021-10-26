package main

import (
	"fmt"
	. "go-leetcode/src/datastruct"
)

var (
	list = [][]int{}
	ints = []int{}
	target = 0
	sum = 0
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	sum = 0
	list = [][]int{}
	ints = []int{}
	target = targetSum
	if root == nil {
		return list
	}
	dfs(root)
	return list
}

func dfs(root *TreeNode)  {
	//添加当前节点到和中
	sum += root.Val
	ints = append(ints,root.Val)
	//到达叶子节点
	if root.Left == nil && root.Right == nil {
		if sum == target {
			fmt.Printf("sum = %v ints = %v\n",sum,ints)
			cInts := make([]int, len(ints))
			copy(cInts,ints)
			fmt.Printf("cInts = %v\n",cInts)
			list = append(list,cInts)
		}
	}
	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
	//移除当前节点（回溯）
	sum -= root.Val
	ints = ints[:len(ints)-1]
}
