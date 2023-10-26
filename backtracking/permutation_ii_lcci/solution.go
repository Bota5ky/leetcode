package permutation_ii_lcci

// 面试题 08.08. 有重复字符串的排列组合 https://leetcode.cn/problems/permutation-ii-lcci/
func permutation2(S string) []string {
	if len(S) == 1 {
		return []string{S}
	}
	var res []string
	m := make(map[string]bool)
	for i := len(S) - 1; i >= 0; i-- {
		rest := S[:i] + S[i+1:]
		strings := permutation2(rest)
		for j := 0; j < len(strings); j++ {
			strings[j] = strings[j] + string(S[i])
			if m[strings[j]] == false {
				res = append(res, strings[j])
				m[strings[j]] = true
			}
		}
	}
	return res
}
