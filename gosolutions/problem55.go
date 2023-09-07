package gosolutions

// https://leetcode-cn.com/problems/jump-game/
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	p := 0
	for {
		max := 0
		if p+nums[p] >= len(nums)-1 {
			return true
		}
		if nums[p] == 0 {
			break
		}
		for i := p; i < len(nums) && i <= p+nums[p]; i++ {
			if i+nums[i] >= max+nums[max] {
				max = i
			}
		}
		p = max
	}
	return false
}
