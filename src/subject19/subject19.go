package main

import (
	. "go-leetcode/src/datastruct"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//为head节点添加新的头节点
	resultNode := &ListNode{Val: -1,Next: head}
	nodes := []*ListNode{}
	nodes = append(nodes,resultNode)
	for head != nil {
		nodes = append(nodes,head)
		head = head.Next
	}
	size := len(nodes)
	if n == 1 { //如果n=1代表删除最后一个节点，则直接置倒数第二个节点的next为nil
		nodes[size-n-1].Next = nil
	} else {
		nodes[size-n-1].Next = nodes[size-n+1]
	}
	return resultNode.Next
}
