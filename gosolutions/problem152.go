package gosolutions

// https://leetcode-cn.com/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	premax, premin, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			premax, premin = premin, premax
		}
		premax = max(premax*nums[i], nums[i])
		premin = min(premin*nums[i], nums[i])
		res = max(res, premax)
	}
	return res
}
