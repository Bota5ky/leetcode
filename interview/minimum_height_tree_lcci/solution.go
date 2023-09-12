package minimum_height_tree_lcci

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

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
