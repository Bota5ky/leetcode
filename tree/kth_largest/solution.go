package er_cha_sou_suo_shu_de_di_kda_jie_dian_lcof

import . "leetcode/_model"

// LCR 174. 寻找二叉搜索树中的目标节点 https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
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
