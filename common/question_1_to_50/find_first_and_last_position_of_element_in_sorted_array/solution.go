package find_first_and_last_position_of_element_in_sorted_array

// 34. 在排序数组中查找元素的第一个和最后一个位置 https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
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
