package longest_consecutive_sequence

// 128. 最长连续序列 https://leetcode.cn/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if m[nums[i]-1] == false {
			cnt := 0
			for j := nums[i]; m[j]; j++ {
				cnt++
			}
			if cnt > max {
				max = cnt
			}
		}
	}
	return max
}
