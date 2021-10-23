package datastruct

import "fmt"

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

type ListNode struct {
	Val int
	Next *ListNode
}

func PrintListNode(listNode *ListNode)  {
	fmt.Print("[")
	for listNode != nil {
		fmt.Print(" ",listNode.Val)
		listNode = listNode.Next
	}
	fmt.Print(" ]")
	fmt.Println()
}