package legal_binary_search_tree_lcci

import . "leetcode/model"

// https://leetcode.cn/problems/legal-binary-search-tree-lcci/
func isValidBST(root *TreeNode) bool {
	var vals []int
	midScan(root, &vals)
	for i := 1; i < len(vals); i++ {
		if vals[i] <= vals[i-1] {
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
