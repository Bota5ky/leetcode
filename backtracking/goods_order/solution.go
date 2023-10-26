package zi_fu_chuan_de_pai_lie_lcof

// LCR 157. 套餐内商品的排列顺序 https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/
func goodsOrder(s string) []string {
	if len(s) < 2 {
		return []string{s}
	}
	strings := goodsOrder(s[1:])
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
