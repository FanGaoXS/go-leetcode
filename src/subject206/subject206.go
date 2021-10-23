package main

import (
	. "go-leetcode/src/datastruct"
)

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
	PrintListNode(head)
	result := reverseList(head)
	PrintListNode(result)
}

//迭代旧链表，然后新链表依次添加
func reverseList(head *ListNode) *ListNode {
	//如果链表数量只有0或者1个
	if head == nil || head.Next == nil {
		return head
	}
	list := []*ListNode{}
	for head != nil {
		list = append(list,head)
		head = head.Next
	}
	resultHead := &ListNode{Val: -1}
	tempHead := resultHead	//临时变量存储新链表的头节点的地址
	for i := len(list)-1; i >= 0; i-- {
		resultHead.Next = list[i]
		resultHead = resultHead.Next
	}
	resultHead.Next = nil
	return tempHead.Next
}

//头插法

func reverseList2(head *ListNode) *ListNode {
	//如果链表数量只有0或者1个
	if head == nil || head.Next == nil {
		return head
	}
	return nil
}


