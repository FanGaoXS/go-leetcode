package main

import (
	. "go-leetcode/src/datastruct"
	"sort"
)

func main() {



}

func mergeKLists(lists []*ListNode) *ListNode {
	ints := []int{}
	//依次迭代所有链表中的节点并且添加到新的数组中
	for i := 0; i < len(lists); i++ {
		for lists[i] != nil {
			ints = append(ints,lists[i].Val);
			lists[i] = lists[i].Next
		}
	}
	//对节点的值进行排序
	sort.Ints(ints)
	//创建新的链表
	resultNode := &ListNode{Val: -1}
	tempNode := resultNode
	//将排序好的节点的值创建新的链表
	for i := 0; i < len(ints); i++ {
		resultNode.Next = &ListNode{Val: ints[i]}
		resultNode = resultNode.Next
	}
	return tempNode.Next
}
