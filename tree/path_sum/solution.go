package path_sum

import . "leetcode/_model"

// 112. 路径总和 https://leetcode.cn/problems/path-sum/
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
