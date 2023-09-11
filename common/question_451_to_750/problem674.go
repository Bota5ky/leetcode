package question_451_to_750

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, cnt := 1, 1
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] > nums[i-1] {
			cnt++
			if cnt > max {
				max = cnt
			}
		} else {
			cnt = 1
		}
	}
	return max
}
