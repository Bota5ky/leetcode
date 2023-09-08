package common

import "math"

// https://leetcode-cn.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// func isValidBST(root *TreeNode) bool {
// 	var nums []int
// 	preOrder(root, &nums)
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] <= nums[i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }
// func preOrder(root *TreeNode, nums *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	preOrder(root.Left, nums)
// 	*nums = append(*nums, root.Val)
// 	preOrder(root.Right, nums)
// }
