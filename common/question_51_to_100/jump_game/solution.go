package jump_game

// 55. 跳跃游戏 https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	p := 0
	for {
		maximum := 0
		if p+nums[p] >= len(nums)-1 {
			return true
		}
		if nums[p] == 0 {
			break
		}
		for i := p; i < len(nums) && i <= p+nums[p]; i++ {
			if i+nums[i] >= maximum+nums[maximum] {
				maximum = i
			}
		}
		p = maximum
	}
	return false
}
