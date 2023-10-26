package kth_smallest_element_in_a_bst

import . "leetcode/_model"

// 230. 二叉搜索树中第K小的元素 https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	var nums []int
	traverse(root, &nums)
	return nums[k-1]
}

func traverse(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, nums)
	*nums = append(*nums, root.Val)
	traverse(root.Right, nums)
}

// 迭代 0(H+k)复杂度
func kthSmallest3(root *TreeNode, k int) int {
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			k--
			if k == 0 {
				return cur.Val
			}
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}
	return -1
}
