package main

import (
	. "go-leetcode/src/datastruct"
)


func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		ints := []int{}
		for i := 0; i < size; i++ {
			poll := queue[0]
			queue = queue[1:len(queue)]	//完成出队
			ints = append(ints, poll.Val)
			if poll.Left != nil {
				queue = append(queue,poll.Left)
			}
			if poll.Right != nil {
				queue = append(queue,poll.Right)
			}
		}
		result = append(result,ints)
	}
	return result
}