package remove_duplicates_from_sorted_array

// 26. 删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates2(nums []int) int {
	var i, j int
	for i, j = 0, 1; j < len(nums); {
		if nums[j] == nums[i] {
			j++
			continue
		} else {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}
