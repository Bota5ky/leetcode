package path_sum

// https://leetcode.cn/problems/path-sum/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	diff := sum - root.Val
	if root.Left == nil && root.Right == nil {
		return diff == 0
	}

	return hasPathSum(root.Left, diff) || hasPathSum(root.Right, diff)
}

// func hasPathSum(root *TreeNode, sum int) bool {
// 	if root == nil {
// 		return false
// 	}
// 	return dfs2(root, 0, sum)
// }

// func dfs2(root *TreeNode, preSum int, target int) bool {
// 	if root.Left == nil && root.Right == nil && preSum+root.Val == target {
// 		return true
// 	}
// 	var a, b bool
// 	if root.Left != nil {
// 		a = dfs2(root.Left, preSum+root.Val, target)
// 	}
// 	if root.Right != nil {
// 		b = dfs2(root.Right, preSum+root.Val, target)
// 	}
// 	return a || b
// }
