package shu_zu_zhong_chu_xian_ci_shu_chao_guo_yi_ban_de_shu_zi_lcof

import "sort"

// 剑指 Offer 39. 数组中出现次数超过一半的数字 https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// 面试题 17.10. 主要元素 https://leetcode.cn/problems/find-majority-element-lcci/
// 169. 多数元素 https://leetcode.cn/problems/majority-element/
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
