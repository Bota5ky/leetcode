package leetcode

//https://leetcode-cn.com/problems/permutation-ii-lcci/
func permutation2(S string) []string {
	if len(S) == 1 {
		return []string{S}
	}
	var res []string
	m := make(map[string]bool)
	for i := len(S) - 1; i >= 0; i-- {
		rest := S[:i] + S[i+1:]
		strs := permutation2(rest)
		for j := 0; j < len(strs); j++ {
			strs[j] = strs[j] + string(S[i])
			if m[strs[j]] == false {
				res = append(res, strs[j])
				m[strs[j]] = true
			}
		}
	}
	return res
}
