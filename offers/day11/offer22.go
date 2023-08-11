package day11

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	slow := head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	nodes := make([]*ListNode, 0)

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	return nodes[len(nodes)-k]
}
