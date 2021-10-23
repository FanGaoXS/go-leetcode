package main

import . "go-leetcode/src/datastruct"

func main() {

}

//set
func hasCycle(head *ListNode) bool {
	//如果链表节点数量为0或1个
	if head == nil || head.Next == nil {return false}
	var set = make(map[*ListNode]bool)
	for head != nil {
		if set[head] {return true}	//如果head已经存在于链表中说明有环
		set[head] = true
		head = head.Next
	}
	return false
}

//快慢指针
func hasCycle2(head *ListNode) bool {
	//如果链表节点数量为0或1个
	if head == nil || head.Next == nil {return false}
	slow := head		//慢指针
	fast := head.Next	//快指针
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
