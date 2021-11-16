package subject104

import (
	. "go-leetcode/src/datastruct"
)

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue,root) //入队
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			poll := queue[0]
			queue = queue[1:] //完成出队
			if poll.Left != nil {
				queue = append(queue,poll.Left)
			}
			if poll.Right != nil {
				queue = append(queue,poll.Right)
			}
		}
		level++
	}
	return level
}
