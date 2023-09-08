package question_151_to_200

import "leetcode/common/question_101_to_150"

// https://leetcode-cn.com/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	premax, premin, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			premax, premin = premin, premax
		}
		premax = common.max(premax*nums[i], nums[i])
		premin = min(premin*nums[i], nums[i])
		res = common.max(res, premax)
	}
	return res
}
