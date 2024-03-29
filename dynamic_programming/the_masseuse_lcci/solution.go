package the_masseuse_lcci

// 面试题 17.16. 按摩师 https://leetcode.cn/problems/the-masseuse-lcci/
func massage(nums []int) int {
	//dp[i+2]=maxInt(dp[i]+nums[i+2],dp[i-1])
	if len(nums) == 0 {
		return 0
	}
	dp0 := nums[0]
	if len(nums) == 1 {
		return dp0
	}
	dp1 := maxInt(nums[0], nums[1])
	if len(nums) == 2 {
		return dp1
	}
	var dp2 int
	for i := 2; i < len(nums); i++ {
		dp2 = maxInt(dp0+nums[i], dp1)
		dp0 = dp1
		dp1 = dp2
	}
	return dp2
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
