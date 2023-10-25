package binary_tree_right_side_view

import . "leetcode/model"

// 199. 二叉树的右视图 https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		k := len(stack)
		for i := 0; i < k; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		res = append(res, stack[k-1].Val)
		stack = stack[k:]
	}
	return res
}
