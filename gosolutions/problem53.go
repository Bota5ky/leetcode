package gosolutions

// 和offer42相同
// https://leetcode-cn.com/problems/maximum-subarray/
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
