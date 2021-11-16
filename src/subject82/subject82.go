package subject82

import (
	. "go-leetcode/src/datastruct"
	"math"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	resultNode := &ListNode{Val: math.MinInt32}
	tempNode := resultNode
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = head.Next
			continue
		}

		resultNode.Next = head

		resultNode = resultNode.Next
		head = head.Next
	}
	resultNode.Next = nil
	return tempNode.Next
}
