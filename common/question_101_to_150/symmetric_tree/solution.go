package symmetric_tree

// 和offer28相同
// https://leetcode.cn/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	return sym(root, root)
}

func sym(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return sym(a.Left, b.Right) && sym(a.Right, b.Left)
}
