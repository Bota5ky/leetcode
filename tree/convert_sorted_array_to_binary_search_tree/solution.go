package convert_sorted_array_to_binary_search_tree

import . "leetcode/_model"

// 108. 将有序数组转换为二叉搜索树 https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
// 面试题 04.02. 最小高度树 https://leetcode.cn/problems/minimum-height-tree-lcci/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	node := &TreeNode{Val: nums[len(nums)/2]}
	node.Left = sortedArrayToBST(nums[:len(nums)/2])
	node.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	return node
}
