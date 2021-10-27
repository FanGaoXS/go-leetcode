package main

import (
	. "go-leetcode/src/datastruct"
)

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	set := make(map[*ListNode]bool)
	for head != nil {
		if set[head] {
			return head
		}
		set[head] = true
		head = head.Next
	}
	return nil
}
