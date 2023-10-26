package count_complete_tree_nodes

import . "leetcode/_model"

// 222. 完全二叉树的节点个数 https://leetcode.cn/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	var stack []*TreeNode
	sum := 0
	for root != nil || len(stack) > 0 {
		if root != nil {
			sum++
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return sum
}
