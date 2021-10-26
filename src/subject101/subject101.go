package main

import . "go-leetcode/src/datastruct"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root.Left,root.Right);
}

func check(leftNode,rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil{
		return true
	}
	if leftNode == nil || rightNode == nil {
		return false
	}
	if leftNode.Val == rightNode.Val {
		return check(leftNode.Left,rightNode.Right)&&check(leftNode.Right,rightNode.Left);
	}
	return false;
}
