package offer32_3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 标识位来决定正反读node元素
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	res := make([][]int, 0, 0)
	flag := false

	for len(queue) != 0 {
		flag = !flag
		size := len(queue)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if flag {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
