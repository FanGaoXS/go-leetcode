package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func initLinkedList() {
	// 实例化节点
	n1 := &ListNode{Val: 4}
	n2 := &ListNode{Val: 5}
	n3 := &ListNode{Val: 1}
	// 构建指向引用
	n1.Next = n2
	n2.Next = n3
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTree() {
	// 实例化节点
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 1}
	n5 := &TreeNode{Val: 2}
	// 构建指向引用
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
}

func hash() {
	// 初始化hash
	m := make(map[string]int)
	// 建立 key->value 的映射关系
	m["小力"] = 10001
	m["小特"] = 10002
	m["小扣"] = 10003
	// 通过 key 查找 value
	_ = m["小力"] // 10001
	_ = m["小特"] // 10002
	_ = m["小扣"] // 10003
}
