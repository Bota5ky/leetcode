package leetcode

//比offer42严谨
//https://leetcode-cn.com/problems/contiguous-sequence-lcci/
func maxSubArray2(nums []int) int {
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
