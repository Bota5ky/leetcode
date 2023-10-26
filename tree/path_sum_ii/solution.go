package path_sum_ii

import . "leetcode/_model"

// 113. 路径总和 II https://leetcode.cn/problems/path-sum-ii/
// 剑指 Offer 34. 二叉树中和为某一值的路径 https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
func pathSum(root *TreeNode, sum int) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	stack := []*TreeNode{root}
	view := make(map[*TreeNode]int)
	target := root.Val
	temp := []int{target}
	for len(stack) > 0 {
		k := len(stack) - 1
		if view[stack[k]] == 0 {
			if stack[k].Left != nil {
				stack = append(stack, stack[k].Left)
				target += stack[k].Left.Val
				temp = append(temp, stack[k].Left.Val)
			}
			view[stack[k]]++
			continue
		}
		if view[stack[k]] == 1 {
			if stack[k].Right != nil {
				stack = append(stack, stack[k].Right)
				target += stack[k].Right.Val
				temp = append(temp, stack[k].Right.Val)
			}
			view[stack[k]]++
			continue
		}
		if target == sum && stack[k].Left == nil && stack[k].Right == nil {
			var tmp []int
			for i := 0; i < len(temp); i++ {
				tmp = append(tmp, temp[i])
			}
			ret = append(ret, tmp)
		}
		target -= stack[len(stack)-1].Val
		stack = stack[:len(stack)-1]
		temp = temp[:len(temp)-1]
	}
	return ret
}
