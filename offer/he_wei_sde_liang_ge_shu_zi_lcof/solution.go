package offer

// https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/
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
