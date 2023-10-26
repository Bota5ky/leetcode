package binary_tree_zigzag_level_order_traversal

import . "leetcode/_model"

// 103. 二叉树的锯齿形层序遍历 https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var nums [][]int
	if root == nil {
		return nums
	}
	stack := []*TreeNode{root}
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
			if len(nums)%2 == 0 {
				temp = append(temp, stack[i].Val)
			} else {
				temp = append([]int{stack[i].Val}, temp...)
			}
		}
		stack = stack[k:]
		nums = append(nums, temp)
	}
	return nums
}
