package offer

import "sort"

// https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
