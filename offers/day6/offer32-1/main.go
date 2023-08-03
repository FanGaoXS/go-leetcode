package offer32_1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 队列bfs
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		// pop node from the left of the queue
		node := queue[0]
		queue = queue[1:]

		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res
}
