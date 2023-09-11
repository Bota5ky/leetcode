package missing_number_lcci

// https://leetcode.cn/problems/missing-number-lcci/
func missingNumber2(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}
