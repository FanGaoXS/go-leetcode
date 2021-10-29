package main

import (
	"fmt"
	. "go-leetcode/src/datastruct"
	"math"
)

func main() {
	var ints = make([]int,5)
	for i := range ints {
		fmt.Println("ints[i] =",ints[i])
	}
}

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{}
	queue = append(queue,root)
	level := 1
	for len(queue) != 0{
		size := len(queue)
		nextIdealSize := int(math.Pow(2, float64(level)))
		var nextBools = make([]bool,nextIdealSize)
		for i := 0; i < size; i++ {
			//完成出队
			poll := queue[0]
			queue = queue[1:]
			if poll.Left != nil {
				queue = append(queue,poll.Left)
				nextBools[i*2] = true
			}
			if poll.Right != nil {
				queue = append(queue,poll.Right)
				nextBools[i*2+1] = true
			}
		}
		//此时的队列的大小是下层的元素个数
		if len(queue) != nextIdealSize {
			for i := 0; i < len(nextBools)-1; i++ {
				if !nextBools[i]&&nextBools[i+1] {
					return false
				}
			}
			for len(queue) != 0{
				poll := queue[0]
				queue = queue[1:]
				if poll.Left != nil || poll.Right != nil {
					return false
				}
			}
		}
		level++
	}
	return true
}
