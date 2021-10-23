package main

import (
 	. "go-leetcode/src/datastruct"
)

func main() {
	listNode := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	PrintListNode(&listNode)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visitedSet := map[*ListNode]bool{} //利用map实现set
	//fmt.Printf("%v\n",headA)
	for headA!=nil {	//迭代headA并一次添加到set中
		//fmt.Println(headA.Val)
		visitedSet[headA] = true
		headA = headA.Next
	}
	for headB!=nil {
		//fmt.Println(headB.Val)
		if visitedSet[headB] {//如果headB存在于set中，说明相交
			return headB
		}
		headB = headB.Next
	}
	return nil
}
