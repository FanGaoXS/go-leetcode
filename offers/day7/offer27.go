package day7

// 递归：交换左右节点
func mirrorTree(root *TreeNode) *TreeNode {
	swap(root)
	return root
}

func swap(node *TreeNode) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	swap(node.Left)
	swap(node.Right)
}

// 栈：交换所有节点的左右节点
func mirrorTree2(root *TreeNode) *TreeNode {
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
