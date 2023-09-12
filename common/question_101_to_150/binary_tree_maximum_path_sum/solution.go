package binary_tree_maximum_path_sum

import . "leetcode/model"

// 124. 二叉树中的最大路径和 https://leetcode.cn/problems/binary-tree-maximum-path-sum/
func maxPathSum(root *TreeNode) int {
	m := root.Val
	maxWeight(root, &m)
	return m
}

func maxWeight(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}
	//包含根节点的权重
	left := max(maxWeight(root.Left, m), 0)
	right := max(maxWeight(root.Right, m), 0)
	*m = max(*m, left+right+root.Val)
	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
