package main

import (
	. "go-leetcode/src/datastruct"
	"math"
)

func isValidBST(root *TreeNode) bool {
	return dfs(root,math.MinInt64,math.MaxInt64)
}

func dfs(root *TreeNode,lower int64,upper int64) bool {
	if root == nil {
		return true
	}
	if int64(root.Val) <= lower || int64(root.Val) >= upper {
		return false
	}
	return dfs(root.Left,lower,int64(root.Val)) && dfs(root.Right,int64(root.Val),upper)
}

