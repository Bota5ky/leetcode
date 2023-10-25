package jump_game_ii

// 45. 跳跃游戏 II https://leetcode.cn/problems/jump-game-ii/
func jump(nums []int) int {
	target := len(nums) - 1
	stack := []int{0, 0}
	cnt := 0
	for stack[1] < target {
		max := -1
		for i := stack[0]; i <= stack[1]; i++ {
			if i+nums[i] > max {
				max = i + nums[i]
			}
		}
		stack[0] = stack[1] + 1
		stack[1] = max
		cnt++
	}
	return cnt
}
