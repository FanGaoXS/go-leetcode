package main

import . "go-leetcode/src/datastruct"
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil || left == right {
		return head
	}
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes,head)
		head = head.Next;
	}
	resultNode := &ListNode{Val: -1}
	tempNode := resultNode
	for i := 0; i < len(nodes); i++ {
		if i>=left-1 && i<=right-1 {
			resultNode.Next = nodes[right-1-(i-(left-1))]
		} else {
			resultNode.Next = nodes[i]
		}
		resultNode = resultNode.Next
	}
	resultNode.Next = nil
	return tempNode.Next
}
