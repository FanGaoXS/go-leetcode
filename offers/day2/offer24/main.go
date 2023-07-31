package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	count := 0
	values := make([]int, 0, 0)
	for head != nil {
		values = append(values, head.Val)

		head = head.Next
		count++
	}
	// now head is nil
	head = &ListNode{Val: -1}
	tempHead := head
	for i := count - 1; i >= 0; i-- {
		head.Next = &ListNode{Val: values[i]}
		head = head.Next
	}

	return tempHead.Next
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var resNode *ListNode
	for head != nil {
		tempNode := head.Next

		head.Next = resNode
		resNode = head

		head = tempNode
	}

	return resNode
}

func reverseList3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var newHead *ListNode
	for head != nil {
		newHead = &ListNode{Val: head.Val, Next: newHead}
		head = head.Next
	}
	return newHead
}
