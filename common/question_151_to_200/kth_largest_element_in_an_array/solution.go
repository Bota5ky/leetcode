package question_151_to_200

// https://leetcode.cn/problems/kth-largest-element-in-an-array/
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
