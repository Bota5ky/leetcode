package kth_largest_element_in_an_array

// 215. 数组中的第K个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	for j := 1; j <= k; j++ {
		for i := 0; i < len(nums)-j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums[len(nums)-k]
}
