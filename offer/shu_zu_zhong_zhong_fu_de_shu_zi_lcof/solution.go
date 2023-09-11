package offer

// 剑指 Offer 03. 数组中重复的数字 https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			p := nums[i]
			if nums[i] != nums[p] {
				nums[i], nums[p] = nums[p], nums[i]
			} else {
				return nums[i]
			}
		}
	}
	return -1
}
