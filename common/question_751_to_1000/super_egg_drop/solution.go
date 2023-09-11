package question_751_to_1000

// 887. 鸡蛋掉落 https://leetcode.cn/problems/super-egg-drop/
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