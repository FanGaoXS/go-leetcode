package main

import . "go-leetcode/src/datastruct"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes,head)
		head = head.Next;
	}
	size := len(nodes)
	if len(nodes)%2 != 0 {
		size = len(nodes)-1
	}
	resultNode := &ListNode{Val: -1}
	tempNode := resultNode
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			resultNode.Next = nodes[i+1]
		} else {
			resultNode.Next = nodes[i-1]
		}
		resultNode = resultNode.Next
	}
	if len(nodes)%2 != 0 {
		resultNode.Next = nodes[len(nodes)-1]
		resultNode = resultNode.Next
	}
	resultNode.Next = nil
	return tempNode.Next
}
