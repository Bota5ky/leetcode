package product_of_array_except_self

// 238. 除自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/
// LCR 191. 按规则计算统计结果 https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	temp := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		res[i] *= temp
		temp *= nums[i]
	}
	return res
}
