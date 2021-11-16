package subject104

import (
	. "go-leetcode/src/datastruct"
	"math"
)

var (
	result int
	depth int
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result = math.MinInt32
	depth = 1
	dfs(root)
	return result
}

func dfs(root *TreeNode)  {
	if root.Left == nil && root.Right == nil {
		if depth > result {
			result = depth
		}
		return
	}
	depth ++
	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
	depth --
}
