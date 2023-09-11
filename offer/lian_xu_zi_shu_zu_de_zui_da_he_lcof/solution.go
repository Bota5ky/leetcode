package offer

// 剑指 Offer 42. 连续子数组的最大和 https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func maxSubArray(nums []int) int {
	sum := 0
	max := -111
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
