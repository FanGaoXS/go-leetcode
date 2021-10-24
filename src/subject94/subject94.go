package main

import . "go-leetcode/src/datastruct"

var result = []int{}

func inorderTraversal(root *TreeNode) []int {
	//注意在leetcode的网页ide中利用golang使用全局变量需要先清空，不然会获取到上次的result（bug）
	result = []int{}
	if root == nil {
		return result
	}
	dfs(root)
	return result
}

func dfs(root *TreeNode)  {
	if root == nil {
		return
	}
	dfs(root.Left)		//遍历左子树
	result = append(result,root.Val)
	dfs(root.Right)		//遍历右子树
}

