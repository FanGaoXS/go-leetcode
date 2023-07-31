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
func reversePrint(head *ListNode) []int {
	count := 0
	values := make([]int, 0, 0)
	for head != nil {
		values = append(values, head.Val)

		count++
		head = head.Next
	}
	res := make([]int, 0, count)
	for i := count - 1; i >= 0; i-- {
		res = append(res, values[i])
	}

	return res
}
