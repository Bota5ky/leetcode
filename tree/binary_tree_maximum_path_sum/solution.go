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
	left := maxInt(maxWeight(root.Left, m), 0)
	right := maxInt(maxWeight(root.Right, m), 0)
	*m = maxInt(*m, left+right+root.Val)
	return root.Val + maxInt(left, right)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
