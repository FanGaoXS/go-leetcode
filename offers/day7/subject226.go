package day7

func invert(node *TreeNode) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	invert(node.Left)
	invert(node.Right)
}

// 递归：交换左右节点
func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}

		n.Left, n.Right = n.Right, n.Left
	}

	return root
}
