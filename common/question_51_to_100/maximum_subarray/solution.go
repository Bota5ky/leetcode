package maximum_subarray

// 53. 最大子数组和 https://leetcode.cn/problems/maximum-subarray/
// 剑指 Offer 42. 连续子数组的最大和 https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func maxSubArray(nums []int) int {
	sum, max := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
