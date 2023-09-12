package find_majority_element_lcci

// 面试题 17.10. 主要元素 https://leetcode.cn/problems/find-majority-element-lcci/
// 剑指 Offer 39. 数组中出现次数超过一半的数字 https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// 169. 多数元素 https://leetcode.cn/problems/majority-element/
func majorityElement2(nums []int) int {
	cnt := 0
	cmp := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			cnt++
			cmp = nums[i]
		} else if cmp == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	if cnt > 0 {
		return cmp
	}
	return -1
}
