package day7

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if dfs(A, B) {
		return true
	}
	if isSubStructure(A.Left, B) {
		return true
	}
	if isSubStructure(A.Right, B) {
		return true
	}

	return false
}

func dfs(a, b *TreeNode) bool {
	// a包含b
	if b == nil {
		return true
	}
	// a大于b
	if a == nil {
		return false
	}
	if a.Val == b.Val {
		return dfs(a.Left, b.Left) && dfs(a.Right, b.Right)
	}
	return false
}
