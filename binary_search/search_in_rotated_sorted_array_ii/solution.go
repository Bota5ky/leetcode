package search_in_rotated_sorted_array_ii

// 81. 搜索旋转排序数组 II https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target { //mid值等于target直接返回
			return true
		}
		if nums[m] > nums[l] { //mid在旋转点前面
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[l] {
			if target <= nums[r] && target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			l++
		}
	}
	return false
}
