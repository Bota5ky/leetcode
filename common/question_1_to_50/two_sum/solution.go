package two_sum

// 1. 两数之和 https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for c, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, c}
		}
		m[v] = c
	}
	return []int{}
}
