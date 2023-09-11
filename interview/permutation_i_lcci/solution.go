package permutation_i_lcci

// https://leetcode.cn/problems/permutation-i-lcci/
func permutation(S string) []string {
	var ret []string
	if len(S) < 2 {
		return []string{S}
	}
	for i := 0; i < len(S); i++ {
		single := S[i]
		temp := permutation(S[:i] + S[i+1:])
		for j := 0; j < len(temp); j++ {
			ret = append(ret, string(single)+temp[j])
		}
	}
	return ret
}
