package zai_pai_xu_shu_zu_zhong_cha_zhao_shu_zi_lcof

// LCR 172. 统计目标成绩的出现次数 https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func countTarget(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	//找左边
	for i < j {
		mid := (i + j) / 2
		if nums[mid] >= target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	left := i
	if j < 0 || nums[left] != target {
		return 0
	}
	//找右边
	i = 0
	j = len(nums) - 1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] > target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	right := i
	if nums[right] == target {
		right++
	}
	return right - left
}
