package lian_xu_zi_shu_zu_de_zui_da_he_lcof

// 剑指 Offer 42. 连续子数组的最大和 https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
// 53. 最大子数组和 https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	sum := 0
	maximum := -111
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if maximum < sum {
			maximum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maximum
}
