package day7

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return compare(root.Left, root.Right)
}

// 对称二叉树定义： 对于树中 任意两个对称节点 L 和 R ，一定有：
//
//L.val = R.val ：即此两对称节点值相等。
//L.left.val = R.right.val ：即 L 的 左子节点 和 R 的 右子节点 对称；
//L.right.val = R.left.val ：即 L 的 右子节点 和 R 的 左子节点 对称。

func compare(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && compare(l.Left, r.Right) && compare(l.Right, r.Left)
}
