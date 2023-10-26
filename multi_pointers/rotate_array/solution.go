package rotate_array

// 189. 轮转数组 https://leetcode.cn/problems/rotate-array/
func rotate(nums []int, k int) {
	k = k % len(nums)
	temp := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, temp)
}
