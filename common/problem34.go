package common

// 和offer53-I类似
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	//找左边
	for i < j {
		mid := (i + j) / 2
		if nums[mid] >= target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	left := i
	if j < 0 || nums[left] != target {
		return []int{-1, -1}
	}
	//找右边
	i = 0
	j = len(nums) - 1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] > target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	right := i
	if nums[right] == target {
		right++
	}
	return []int{left, right - 1}
}
