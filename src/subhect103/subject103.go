package main

import (
	. "go-leetcode/src/datastruct"
)

func main() {

}

func zigzagLevelOrder(root *TreeNode) [][]int {
	list := [][]int{}
	if root == nil { return list }
	queue := []*TreeNode{}
	queue = append(queue,root)
	flag := true
	for len(queue) != 0 {
		size := len(queue)
		ints := make([]int,size)
		for i := 0; i < size; i++ {
			poll := queue[0]
			queue = queue[1:len(queue)]	//完成出队
			if flag {
				//正向添加
				ints[i] = poll.Val
			} else {
				//反向添加
				ints[size-1-i] = poll.Val
			}
			if poll.Left != nil {
				queue = append(queue,poll.Left)
			}
			if poll.Right != nil {
				queue = append(queue,poll.Right)
			}

		}
		list = append(list,ints)
		flag = !flag //该层遍历完，取反
	}
	return list
}
