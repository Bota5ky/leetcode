package leetcode

//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
func kthLargest(root *TreeNode, k int) int {
	var nums []int
	list(root, &nums)
	return nums[len(nums)-k]
}

//中序遍历
func list(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	list(root.Left, nums)
	*nums = append(*nums, root.Val)
	list(root.Right, nums)
}
