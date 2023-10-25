package word_break

// 139. 单词拆分 https://leetcode.cn/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for _, v := range wordDict {
		m[v] = true
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && m[s[j:i+1]] {
				dp[i+1] = true
			}
		}
	}
	return dp[len(s)]
}
