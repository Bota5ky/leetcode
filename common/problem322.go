package common

import "math"

// DP问题
// https://leetcode-cn.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if len(coins) == 0 {
		return -1
	}
	//dp[i]=min(dp[i-coins])+1 遍历coins
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		//取最小值
		min := math.MaxInt32
		for _, val := range coins {
			if i-val >= 0 && dp[i-val] >= 0 && dp[i-val] < min {
				min = dp[i-val]
			}
		}
		//
		if min == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = min + 1
		}
	}
	return dp[amount]
}
