package main

import . "go-leetcode/src/datastruct"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}
	head.PrintListNode()
	reorderList(head)
	head.PrintListNode()
}

func reorderList(head *ListNode)  {
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes,head)
		head = head.Next
	}
	flag := true
	l := 1
	r := len(nodes)-1
	tempNode := nodes[0]
	for i := 1; i < len(nodes); i++ {
		if flag {
			tempNode.Next = nodes[r]
			r--
		} else {
			tempNode.Next = nodes[l]
			l++
		}
		flag = !flag
		tempNode = tempNode.Next
	}
	tempNode.Next = nil
	head = nodes[0]
}
