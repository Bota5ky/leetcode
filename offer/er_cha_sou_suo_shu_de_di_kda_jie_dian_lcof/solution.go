package offer

import . "leetcode/model"

// 剑指 Offer 54. 二叉搜索树的第k大节点 https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
func kthLargest(root *TreeNode, k int) int {
	var nums []int
	list(root, &nums)
	return nums[len(nums)-k]
}

// 中序遍历
func list(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	list(root.Left, nums)
	*nums = append(*nums, root.Val)
	list(root.Right, nums)
}
