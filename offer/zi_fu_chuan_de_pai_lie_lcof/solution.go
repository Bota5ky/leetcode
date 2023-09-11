package zi_fu_chuan_de_pai_lie_lcof

// 剑指 Offer 38. 字符串的排列 https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/
func permutation(s string) []string {
	if len(s) < 2 {
		return []string{s}
	}
	strings := permutation(s[1:])
	var ret []string
	m := make(map[string]bool)
	for i := 0; i < len(strings); i++ {
		temp := []byte{s[0]}
		temp = append(temp, []byte(strings[i])...)
		for j := 0; j < len(temp); j++ {
			t := make([]byte, len(temp))
			copy(t, temp)
			t[0], t[j] = t[j], t[0]
			if m[string(t)] == false {
				ret = append(ret, string(t))
				m[string(t)] = true
			}
		}
	}
	return ret
}
