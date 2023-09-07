package leetcode

//https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func search2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			if nums[m] >= nums[0] || target < nums[0] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target >= nums[0] || nums[m] < nums[0] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
