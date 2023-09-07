package leetcode

//返回值    返回[1:]开始后面的所有组合结果
//每次做什么
//终止条件
//https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
func permutation(s string) []string {
	if len(s) < 2 {
		return []string{s}
	} //  []string{a}  "a"
	strs := permutation(s[1:])
	var ret []string
	m := make(map[string]bool)
	for i := 0; i < len(strs); i++ { //strs[i]
		temp := []byte{s[0]}
		temp = append(temp, []byte(strs[i])...)
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
