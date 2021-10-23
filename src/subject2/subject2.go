package main

import . "go-leetcode/src/datastruct"

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := &ListNode{Val: -1}
	tempNode := resultNode
	add := false
	for add || l1 != nil || l2 != nil {
		i := 0
		j := 0
		sum := 0
		if l1 != nil {
			i = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			j = l2.Val
			l2 = l2.Next
		}
		if add {
			sum = i + j + 1
		} else {
			sum = i + j
		}
		if sum >= 10 {
			sum = sum%10
			add = true
		} else {
			add = false
		}
		resultNode.Next = &ListNode{Val: sum}
		resultNode = resultNode.Next
	}
	return tempNode.Next
}
