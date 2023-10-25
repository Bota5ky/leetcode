package maximum_subarray

// 53. 最大子数组和 https://leetcode.cn/problems/maximum-subarray/
// 面试题 16.17. 连续数列 https://leetcode.cn/problems/contiguous-sequence-lcci/ 更严谨
// 剑指 Offer 42. 连续子数组的最大和 https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if maxSum < nums[i] {
			maxSum = nums[i]
		}
	}
	return maxSum
}
