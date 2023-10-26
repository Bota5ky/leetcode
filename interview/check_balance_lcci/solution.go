package check_balance_lcci

import . "leetcode/_model"

// 面试题 04.04. 检查平衡性 https://leetcode.cn/problems/check-balance-lcci/
// 剑指 Offer 55 - II. 平衡二叉树 https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
// 110. 平衡二叉树 https://leetcode.cn/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	var depth int
	return isBalancedTree(root, &depth)
}
func isBalancedTree(root *TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}
	var left, right int
	if isBalancedTree(root.Left, &left) && isBalancedTree(root.Right, &right) {
		if left >= right {
			*depth = 1 + left
		} else {
			*depth = 1 + right
		}
		if left-right <= 1 && left-right >= -1 {
			return true
		}
	}
	return false
}
