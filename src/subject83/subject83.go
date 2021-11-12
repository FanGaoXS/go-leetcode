package subject83

import . "go-leetcode/src/datastruct"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tempNode := head
	for head != nil {
		//如果下一节点存在并且当前节点的值和下一节点的值相同
		//则当前节点指向下下一节点
		if head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue	//当前节点继续与下一节点比较
		}
		head = head.Next
	}
	return tempNode
}

