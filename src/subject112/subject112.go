package main

import (
	. "go-leetcode/src/datastruct"
)

var (
	sum = 0
	target = 0
	result = false
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	sum = 0
	result = false
	target = targetSum
	dfs(root)
	return result
}

func dfs(root *TreeNode) {
	if result {
		return
	}
	sum+=root.Val
	//没有左右子树了（到达叶子节点）
	if root.Left == nil && root.Right == nil {
		if sum == target {
			result = true
		}
		sum-=root.Val
		return
	}
	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
	sum-=root.Val
}
