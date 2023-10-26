package shu_zu_zhong_zhong_fu_de_shu_zi_lcof

// LCR 120. 寻找文件副本 https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
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
