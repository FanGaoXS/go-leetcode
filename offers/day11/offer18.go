package day11

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && head.Val == val {
		return nil
	}
	if head.Next != nil && head.Val == val {
		return &ListNode{
			Val:  head.Next.Val,
			Next: head.Next.Next,
		}
	}

	temp := head

	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}

	return temp
}

func deleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	node := &ListNode{Val: -1}
	temp := node
	for head != nil {
		if head.Val != val {
			node.Next = &ListNode{Val: head.Val}
			node = node.Next
		}
		head = head.Next
	}

	return temp.Next
}
