package leetcode

//https://leetcode-cn.com/problems/minimum-cost-for-tickets/
func mincostTickets(days []int, costs []int) int {
	//1 7 30
	//dp[i]=min(dp[i-1]+costs[0],dp[i-7]+costs[1],dp[i-30]+costs[2])
	dp := make([]int, days[len(days)-1]+1)
	for i, j := 1, 0; i < len(dp); i++ {
		dp[i] = minDp(dp, costs, i, days[j])
		if days[j] == i {
			j++
		}
	}
	return dp[len(dp)-1]
}

func minDp(dp, costs []int, i, target int) int {
	min := dp[i-1] + costs[0]
	if target != i {
		min = dp[i-1]
	}
	for j := i - 2; j >= i-7 && j >= 0; j-- {
		if min > dp[j]+costs[1] {
			min = dp[j] + costs[1]
		}
	}
	for j := i - 8; j >= i-30 && j >= 0; j-- {
		if min > dp[j]+costs[2] {
			min = dp[j] + costs[2]
		}
	}
	return min
}
