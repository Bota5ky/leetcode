package he_wei_sde_liang_ge_shu_zi_lcof

// 剑指 Offer 57. 和为s的两个数字 https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func twoSum(nums []int, target int) []int {
	cnt := make(map[int]bool)
	for _, v := range nums {
		if cnt[target-v] {
			return []int{v, target - v}
		}
		cnt[v] = true
	}
	return []int{}
}
