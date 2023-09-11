package question_751_to_1000

// https://leetcode.cn/problems/super-egg-drop/
func superEggDrop(K int, N int) int {
	dp := make([]int, K+1)
	ans := 0 // 操作的次数
	for dp[K] < N {
		for i := K; i > 0; i-- { // 从后往前计算
			dp[i] = dp[i] + dp[i-1] + 1
		}
		ans++
	}
	return ans
}

/**
for (int m = 1; dp[K][m] < N; m++)
    for (int k = 1; k <= K; k++)
		dp[k][m] = dp[k][m - 1] + dp[k - 1][m - 1] + 1;
https://leetcode.cn/problems/super-egg-drop/solution/ji-ben-dong-tai-gui-hua-jie-fa-by-labuladong/
**/
