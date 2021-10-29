package main

import (
	. "go-leetcode/src/datastruct"
)

func getKthFromEnd(head *ListNode, k int) *ListNode {
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes,head)
		head = head.Next
	}
	return nodes[len(nodes)-k]
}
