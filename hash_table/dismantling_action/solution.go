package di_yi_ge_zhi_chu_xian_yi_ci_de_zi_fu_lcof

import "strings"

// LCR 169. 招式拆解 II https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
func dismantlingAction(s string) byte {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) == 1 {
			return s[i]
		}
	}
	return ' '
}
