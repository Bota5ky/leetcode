package dui_cheng_de_er_cha_shu_lcof

import . "leetcode/model"

// 剑指 Offer 28. 对称的二叉树 https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/
// 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	return isSym(root, root)
}

func isSym(a *TreeNode, b *TreeNode) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return a == nil
	}
	if a.Val != b.Val {
		return false
	}
	return isSym(a.Left, b.Right) && isSym(a.Right, b.Left)
}
