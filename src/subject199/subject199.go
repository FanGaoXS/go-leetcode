package main
import (
	"fmt"
	. "go-leetcode/src/datastruct"
)

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}		//队列
	queue = append(queue,root)	//队列中添加root
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			poll := queue[0]	//取出第一个元素，即先进先出规则
			queue = queue[1:len(queue)]	//截取元素（完成出队）
			if i == size-1 {
				result = append(result,poll.Val)
				fmt.Printf("val = %d\n",poll.Val)
			}
			if poll.Left != nil {
				queue = append(queue,poll.Left)
			}
			if poll.Right != nil {
				queue = append(queue,poll.Right)
			}
		}
	}
	return result
}