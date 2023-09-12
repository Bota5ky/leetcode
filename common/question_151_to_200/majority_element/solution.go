package majority_element

// 169. 多数元素 https://leetcode.cn/problems/majority-element/
// 剑指 Offer 39. 数组中出现次数超过一半的数字 https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// 面试题 17.10. 主要元素 https://leetcode.cn/problems/find-majority-element-lcci/
func majorityElement(nums []int) int {
	temp, cnt := 0, 0
	for _, v := range nums {
		if cnt == 0 || v == temp {
			temp = v
			cnt++
		} else {
			cnt--
		}
	}
	return temp
}
