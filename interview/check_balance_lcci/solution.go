package check_balance_lcci

import . "leetcode/model"

// 后序遍历只需遍历一次,和offer55-II相同
// https://leetcode-cn.com/problems/check-balance-lcci/
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
