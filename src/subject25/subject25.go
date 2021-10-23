package main

import . "go-leetcode/src/datastruct"
func main() {

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	resultNode := &ListNode{Val: -1}
	tempNode := resultNode
	for head != nil {
		nodes := []*ListNode{}
		for i := 0; i < k; i++ {
			if head == nil {
				break
			}
			nodes = append(nodes,head);
			head = head.Next;
		}
		if len(nodes) == k {	//需要反转
			for i := k-1; i >= 0 ; i-- {
				resultNode.Next = nodes[i]
				resultNode = resultNode.Next
			}
		} else {		//不需要反转
			for i := 0; i < len(nodes); i++ {
				resultNode.Next = nodes[i]
				resultNode = resultNode.Next
			}
		}
	}
	resultNode.Next = nil
	return tempNode.Next
}
