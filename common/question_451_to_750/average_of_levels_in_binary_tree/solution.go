package average_of_levels_in_binary_tree

import . "leetcode/model"

// 637. 二叉树的层平均值 https://leetcode.cn/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *TreeNode) []float64 {
	stack := []*TreeNode{root}
	var ret []float64
	for len(stack) > 0 {
		sum := 0
		k := len(stack)
		for i := 0; i < k; i++ {
			sum += stack[i].Val
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		ret = append(ret, float64(sum)/float64(k))
		stack = stack[k:]
	}
	return ret
}
