package house_robber

// 198. 打家劫舍 https://leetcode.cn/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	//dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
