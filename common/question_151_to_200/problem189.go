package question_151_to_200

// https://leetcode-cn.com/problems/rotate-array/
func rotate(nums []int, k int) {
	k = k % len(nums)
	temp := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, temp)
}
