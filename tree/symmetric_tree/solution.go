package symmetric_tree

import . "leetcode/_model"

// 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree/
// LCR 145. 判断对称二叉树 https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/
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
