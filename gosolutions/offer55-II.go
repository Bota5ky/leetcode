package gosolutions

// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depthOfTreeNode(root.Left)
	right := depthOfTreeNode(root.Right)
	if left-right > 1 || left-right < -1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

func depthOfTreeNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depthOfTreeNode(root.Left)
	right := depthOfTreeNode(root.Right)
	if left >= right {
		return 1 + left
	}
	return 1 + right
}

//后序遍历只需遍历一次
// func isBalanced(root *TreeNode) bool {
// 	var depth int
// 	return isBalancedTree(root, &depth)
// }
// func isBalancedTree(root *TreeNode, depth *int) bool {
// 	if root == nil {
// 		*depth = 0
// 		return true
// 	}
// 	var left, right int
// 	if isBalancedTree(root.Left, &left) && isBalancedTree(root.Right, &right) {
// 		if left >= right {
// 			*depth = 1 + left
// 		} else {
// 			*depth = 1 + right
// 		}
// 		if left-right <= 1 && left-right >= -1 {
// 			return true
// 		}
// 	}
// 	return false
// }
