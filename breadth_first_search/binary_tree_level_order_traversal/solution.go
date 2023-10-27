package binary_tree_level_order_traversal

import . "leetcode/_model"

// 102. 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/
// LCR 150. 彩灯装饰记录 II https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func levelOrder(root *TreeNode) [][]int {
	stack := []*TreeNode{root}
	var res [][]int
	if root == nil {
		return res
	}
	for len(stack) > 0 {
		k := len(stack)
		var temp []int
		for i := 0; i < k; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
			temp = append(temp, stack[i].Val)
		}
		res = append(res, temp)
		stack = stack[k:]
	}
	return res
}
