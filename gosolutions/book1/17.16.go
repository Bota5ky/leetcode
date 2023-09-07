package leetcode

//https://leetcode-cn.com/problems/the-masseuse-lcci/
func massage(nums []int) int {
	//dp[i+2]=max(dp[i]+nums[i+2],dp[i-1])
	if len(nums) == 0 {
		return 0
	}
	dp0 := nums[0]
	if len(nums) == 1 {
		return dp0
	}
	dp1 := max(nums[0], nums[1])
	if len(nums) == 2 {
		return dp1
	}
	var dp2 int
	for i := 2; i < len(nums); i++ {
		dp2 = max(dp0+nums[i], dp1)
		dp0 = dp1
		dp1 = dp2
	}
	return dp2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
