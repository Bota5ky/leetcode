package russian_doll_envelopes

import "sort"

// 354. 俄罗斯套娃信封问题 https://leetcode.cn/problems/russian-doll-envelopes/
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	dp := make([]int, len(envelopes))
	dp[0] = 1
	max := 1
	for i := 1; i < len(envelopes); i++ {
		maxDp := 0
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] && dp[j] > maxDp {
				maxDp = dp[j]
			}
		}
		dp[i] = maxDp + 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
