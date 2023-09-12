package move_zeroes

// 283. 移动零 https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
