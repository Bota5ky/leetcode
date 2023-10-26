package legal_binary_search_tree_lcci

import . "leetcode/_model"

// 面试题 04.05. 合法二叉搜索树 https://leetcode.cn/problems/legal-binary-search-tree-lcci/
func isValidBST(root *TreeNode) bool {
	var values []int
	midScan(root, &values)
	for i := 1; i < len(values); i++ {
		if values[i] <= values[i-1] {
			return false
		}
	}
	return true
}

func midScan(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	midScan(root.Left, vals)
	*vals = append(*vals, root.Val)
	midScan(root.Right, vals)
}
