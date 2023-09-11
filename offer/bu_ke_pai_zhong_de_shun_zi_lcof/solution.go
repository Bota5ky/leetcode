package bu_ke_pai_zhong_de_shun_zi_lcof

import "sort"

// 剑指 Offer 61. 扑克牌中的顺子 https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
func isStraight(nums []int) bool {
	sort.Ints(nums)
	numOfZero := 0
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			numOfZero++
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && nums[i] != 0 {
			return false
		}
		if i > 0 && nums[i]-nums[i-1] > 1 && nums[i-1] != 0 {
			numOfZero -= nums[i] - nums[i-1] - 1
		}
	}
	return numOfZero >= 0
}
