package common

// https://leetcode-cn.com/problems/product-of-array-except-self/
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
