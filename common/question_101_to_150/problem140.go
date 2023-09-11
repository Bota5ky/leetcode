package common

// https://leetcode.cn/problems/word-break-ii/
func wordBreak2(s string, wordDict []string) []string {
	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}
	var pre string
	var res []string
	isInDic(s, m, pre, &res)
	return res
}

func isInDic(s string, m map[string]bool, pre string, res *[]string) {
	if canBreak(s, m) == false {
		return
	}
	for i := 1; i <= len(s); i++ {
		if m[s[:i]] {
			if i == len(s) {
				if len(pre) == 0 {
					*res = append(*res, s)
				} else {
					*res = append(*res, pre+" "+s)
				}
			} else {
				if len(pre) == 0 {
					isInDic(s[i:], m, s[:i], res)
				} else {
					isInDic(s[i:], m, pre+" "+s[:i], res)
				}
			}
		}
	}
}

func canBreak(s string, m map[string]bool) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && m[s[j:i+1]] {
				dp[i+1] = true
			}
		}
	}
	return dp[len(s)]
}
