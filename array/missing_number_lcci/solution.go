package missing_number_lcci

// 面试题 17.04. 消失的数字 https://leetcode.cn/problems/missing-number-lcci/
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}
