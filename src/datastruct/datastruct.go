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

func (this *ListNode) PrintListNode()  {
	fmt.Print("[")
	for this != nil {
		fmt.Print(" ",this.Val)
		this = this.Next
	}
	fmt.Print(" ]")
	fmt.Println()
}
