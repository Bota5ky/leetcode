package question_451_to_750

// https://leetcode-cn.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	pre, cnt := 0, 0
	for _, val := range nums {
		pre += val
		if v, ok := m[pre-k]; ok {
			cnt += v
		}
		m[pre]++
	}
	return cnt
}
