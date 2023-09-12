package maximum_product_subarray

// 152. 乘积最大子数组 https://leetcode.cn/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	preMax, preMin, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			preMax, preMin = preMin, preMax
		}
		preMax = max(preMax*nums[i], nums[i])
		preMin = min(preMin*nums[i], nums[i])
		res = max(res, preMax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
